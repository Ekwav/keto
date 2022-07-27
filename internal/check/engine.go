package check

import (
	"context"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/graph"
)

type (
	EngineProvider interface {
		PermissionEngine() *Engine
	}
	Engine struct {
		d EngineDependencies
	}
	EngineDependencies interface {
		relationtuple.ManagerProvider
		config.Provider
		x.LoggerProvider
	}

	// Type aliases for shorter signatures
	RelationTuple = relationtuple.RelationTuple
	Query         = relationtuple.RelationQuery
)

const WildcardRelation = "..."

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

// CheckIsMember checks if the relation tuple's subject has the relation on the
// object in the namespace either directly or indirectly and returns a boolean
// result.
func (e *Engine) CheckIsMember(ctx context.Context, r *RelationTuple, restDepth int) (bool, error) {
	result := e.CheckRelationTuple(ctx, r, restDepth)
	if result.Err != nil {
		return false, result.Err
	}
	return result.Membership == checkgroup.IsMember, nil
}

// CheckRelationTuple checks if the relation tuple's subject has the relation on
// the object in the namespace either directly or indirectly and returns a check
// result.
func (e *Engine) CheckRelationTuple(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.Result {
	// global max-depth takes precedence when it is the lesser or if the request
	// max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}

	resultCh := make(chan checkgroup.Result)
	go e.checkIsAllowed(ctx, r, restDepth)(ctx, resultCh)
	select {
	case result := <-resultCh:
		return result
	case <-ctx.Done():
		return checkgroup.Result{Err: errors.WithStack(ctx.Err())}
	}
}

// checkExpandSubject checks the expansions of the subject set of the tuple.
//
// For a relation tuple n:obj#rel@user, checkExpandSubject first queries for all
// subjects that match n:obj#rel@* (arbirary subjects), and then for each
// subject checks subject@user.
func (e *Engine) checkExpandSubject(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().
			WithFields(r.ToLoggerFields()).
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", r.String()).
			Trace("check expand subject")

		g := checkgroup.New(ctx)

		var (
			subjects  []*RelationTuple
			pageToken string
			err       error
			visited   bool
			innerCtx  = graph.InitVisited(ctx)
			query     = &Query{Namespace: r.Namespace, Object: r.Object, Relation: r.Relation}
		)
		for {
			subjects, pageToken, err = e.d.RelationTupleManager().GetRelationTuples(innerCtx, query, x.WithToken(pageToken))
			if errors.Is(err, herodot.ErrNotFound) {
				g.Add(checkgroup.NotMemberFunc)
				break
			} else if err != nil {
				g.Add(checkgroup.ErrorFunc(err))
				break
			}
			for _, s := range subjects {
				innerCtx, visited = graph.CheckAndAddVisited(innerCtx, s.Subject)
				if visited {
					continue
				}
				if s.Subject.SubjectSet() == nil || s.Subject.SubjectSet().Relation == WildcardRelation {
					continue
				}
				g.Add(e.checkIsAllowed(
					innerCtx,
					&RelationTuple{
						Namespace: s.Subject.SubjectSet().Namespace,
						Object:    s.Subject.SubjectSet().Object,
						Relation:  s.Subject.SubjectSet().Relation,
						Subject:   r.Subject,
					},
					restDepth-1,
				))
			}
			if pageToken == "" || g.Done() {
				break
			}
		}

		resultCh <- g.Result()
	}
}

// checkDirect checks if the relation tuple is in the database directly.
func (e *Engine) checkDirect(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().
			WithField("method", "checkDirect").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", r.String()).
			Trace("check direct")
		if rels, _, err := e.d.RelationTupleManager().GetRelationTuples(ctx, r.ToQuery()); err == nil && len(rels) > 0 {
			resultCh <- checkgroup.Result{
				Membership: checkgroup.IsMember,
				Tree: &expand.Tree{
					Type:  expand.Leaf,
					Tuple: r,
				},
			}
		} else {
			resultCh <- checkgroup.Result{
				Membership: checkgroup.NotMember,
			}
		}
	}
}

// checkIsAllowed checks if the relation tuple is allowed (there is a path from
// the relation tuple subject to the namespace, object and relation) either
// directly (in the database), or through subject-set expansions, or through
// user-set rewrites.
func (e *Engine) checkIsAllowed(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().
			WithField("method", "checkIsAllowed").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", r.String()).
		Trace("check is allowed")

	g := checkgroup.New(ctx)
	g.Add(e.checkDirect(ctx, r, restDepth-1))
	g.Add(e.checkExpandSubject(ctx, r, restDepth))

	relation, err := e.astRelationFor(ctx, r)
	if err != nil {
		g.Add(checkgroup.ErrorFunc(err))
	} else if relation != nil && relation.UsersetRewrite != nil {
		g.Add(e.checkUsersetRewrite(ctx, r, relation.UsersetRewrite, restDepth))
	}

	return g.CheckFunc()
}

func (e *Engine) astRelationFor(ctx context.Context, r *RelationTuple) (*ast.Relation, error) {
	ns, err := e.namespaceFor(ctx, r)
	if err != nil {
		// On an unknown namespace the answer should be "not allowed", not "not
		// found". Therefore we don't return the error here.
		return nil, nil
	}

	// Special case: If Relations is empty, then there is no namespace
	// configuration, and it is not an error that the relation was not found.
	if len(ns.Relations) == 0 {
		return nil, nil
	}

	for _, rel := range ns.Relations {
		if rel.Name == r.Relation {
			return &rel, nil
		}
	}
	return nil, errors.New("relation not found")
}

func (e *Engine) namespaceFor(ctx context.Context, r *RelationTuple) (*namespace.Namespace, error) {
	namespaceManager, err := e.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	ns, err := namespaceManager.GetNamespaceByName(ctx, r.Namespace)
	if err != nil {
		return nil, err
	}
	return ns, nil
}
