// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/ory/keto/internal/x/events"
	"github.com/ory/keto/internal/x/validate"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var _ rts.WriteServiceServer = (*handler)(nil)

func protoTuplesWithAction(deltas []*rts.RelationTupleDelta, action rts.RelationTupleDelta_Action) (filtered []*ketoapi.RelationTuple, err error) {
	for _, d := range deltas {
		if d.Action == action {
			it, err := (&ketoapi.RelationTuple{}).FromDataProvider(&ketoapi.OpenAPITupleData{d.RelationTuple})
			if err != nil {
				return nil, err
			}
			filtered = append(filtered, it)
		}
	}
	return
}

func (h *handler) TransactRelationTuples(ctx context.Context, req *rts.TransactRelationTuplesRequest) (*rts.TransactRelationTuplesResponse, error) {
	events.Add(ctx, h.d, events.RelationtuplesChanged)

	if err := req.ValidateAll(); err != nil {
		return nil, herodot.ErrBadRequest.WithWrap(err).WithReason(err.Error())
	}

	insertTuples, err := protoTuplesWithAction(req.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_INSERT)
	if err != nil {
		return nil, err
	}

	deleteTuples, err := protoTuplesWithAction(req.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_DELETE)
	if err != nil {
		return nil, err
	}

	its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
	if err != nil {
		return nil, err
	}

	err = h.d.RelationTupleManager().TransactRelationTuples(ctx, its[:len(insertTuples)], its[len(insertTuples):])
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))
	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return &rts.TransactRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}

func (h *handler) CreateRelationTuple(ctx context.Context, request *rts.CreateRelationTupleRequest) (*rts.CreateRelationTupleResponse, error) {
	if request.RelationTuple == nil {
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("invalid request: missing relation_tuple"))
	}
	tuple, err := (&ketoapi.RelationTuple{}).FromDataProvider(&ketoapi.OpenAPITupleData{request.RelationTuple})
	if err != nil {
		return nil, err
	}

	if err := tuple.Validate(); err != nil {
		return nil, err
	}

	mapped, err := h.d.Mapper().FromTuple(ctx, tuple)
	if err != nil {
		return nil, err
	}

	if err := h.d.RelationTupleManager().WriteRelationTuples(ctx, mapped...); err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-location", ReadRouteBase+"?"+tuple.ToURLQuery().Encode()))

	return &rts.CreateRelationTupleResponse{RelationTuple: tuple.ToProto()}, nil
}

func (h *handler) DeleteRelationTuples(ctx context.Context, req *rts.DeleteRelationTuplesRequest) (*rts.DeleteRelationTuplesResponse, error) {
	events.Add(ctx, h.d, events.RelationtuplesDeleted)

	var q ketoapi.RelationQuery

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if hasBody := md["hasbody"]; len(hasBody) > 0 && hasBody[0] == "true" {
			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
			return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("body is not allowed for this request"))
		}
	}

	switch {
	case req.RelationQuery != nil:
		q.FromDataProvider(&queryWrapper{req.RelationQuery})
	case req.Query != nil: // nolint
		q.FromDataProvider(&deprecatedQueryWrapper{(*rts.ListRelationTuplesRequest_Query)(req.Query)}) // nolint
	default:
		q.FromDataProvider(&openAPIQueryWrapper{req})
	}

	if q.Namespace == nil || *q.Namespace == "" {
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("Namespace must be set"))
	}

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, &q)
	if err != nil {
		return nil, err
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithError(err.Error()))
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))
	return &rts.DeleteRelationTuplesResponse{}, nil
}

// Create Relationship Request Parameters
//
// swagger:parameters createRelationship
// nolint:deadcode,unused
type createRelationship struct {
	// in: body
	Body createRelationshipBody
}

// Create Relationship Request Body
//
// swagger:model createRelationshipBody
// nolint:deadcode,unused
type createRelationshipBody struct {
	ketoapi.RelationQuery
}

// swagger:route PUT /admin/relation-tuples relationship createRelationship
//
// # Create a Relationship
//
// Use this endpoint to create a relationship.
//
//	Consumes:
//	-  application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  201: relationship
//	  400: errorGeneric
//	  default: errorGeneric
func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	events.Add(ctx, h.d, events.RelationtuplesCreated)

	var rt ketoapi.RelationTuple
	if err := json.NewDecoder(r.Body).Decode(&rt); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithError(err.Error())))
		return
	}

	if err := rt.Validate(); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Logger().WithFields(rt.ToLoggerFields()).Debug("creating relation tuple")

	it, err := h.d.Mapper().FromTuple(ctx, &rt)
	if err != nil {
		h.d.Logger().WithError(err).WithFields(rt.ToLoggerFields()).Errorf("could not map relation tuple to UUIDs")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().WriteRelationTuples(ctx, it...); err != nil {
		h.d.Logger().WithError(err).WithFields(rt.ToLoggerFields()).Errorf("got an error while creating the relation tuple")
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().WriteCreated(w, r,
		ReadRouteBase+"?"+rt.ToURLQuery().Encode(),
		&rt,
	)
}

// swagger:route DELETE /admin/relation-tuples relationship deleteRelationships
//
// # Delete Relationships
//
// Use this endpoint to delete relationships
//
//	Consumes:
//	-  application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  204: emptyResponse
//	  400: errorGeneric
//	  default: errorGeneric
func (h *handler) deleteRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	events.Add(ctx, h.d, events.RelationtuplesDeleted)

	if err := validate.All(r,
		validate.NoExtraQueryParams(ketoapi.RelationQueryKeys...),
		validate.QueryParamsContainsOneOf(ketoapi.NamespaceKey),
		validate.HasEmptyBody(),
	); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	q := r.URL.Query()
	query, err := (&ketoapi.RelationQuery{}).FromURLQuery(q)
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	l := h.d.Logger()
	for k := range q {
		l = l.WithField(k, q.Get(k))
	}
	l.Debug("deleting relationships")

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, query)
	if err != nil {
		h.d.Logger().WithError(err).Errorf("could not map fields to UUIDs")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		l.WithError(err).Errorf("got an error while deleting relationships")
		h.d.Writer().WriteError(w, r, herodot.ErrInternalServerError.WithError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func internalTuplesWithAction(deltas []*ketoapi.PatchDelta, action ketoapi.PatchAction) (filtered []*ketoapi.RelationTuple) {
	for _, d := range deltas {
		if d.Action == action {
			filtered = append(filtered, d.RelationTuple)
		}
	}
	return
}

// swagger:route PATCH /admin/relation-tuples relationship patchRelationships
//
// # Patch Multiple Relationships
//
// Use this endpoint to patch one or more relationships.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  204: emptyResponse
//	  400: errorGeneric
//	  404: errorGeneric
//	  default: errorGeneric
func (h *handler) patchRelationTuples(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	events.Add(ctx, h.d, events.RelationtuplesChanged)

	var deltas []*ketoapi.PatchDelta
	if err := json.NewDecoder(r.Body).Decode(&deltas); err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}
	for _, d := range deltas {
		if d.RelationTuple == nil {
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError("relation_tuple is missing"))
			return
		}
		if err := d.RelationTuple.Validate(); err != nil {
			h.d.Writer().WriteError(w, r, err)
			return
		}

		switch d.Action {
		case ketoapi.ActionInsert, ketoapi.ActionDelete:
		default:
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError("unknown action "+string(d.Action)))
			return
		}
	}

	insertTuples := internalTuplesWithAction(deltas, ketoapi.ActionInsert)
	deleteTuples := internalTuplesWithAction(deltas, ketoapi.ActionDelete)

	its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
	if err != nil {
		h.d.Logger().WithError(err).Errorf("got an error while mapping fields to UUID")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().
		TransactRelationTuples(
			ctx,
			its[:len(insertTuples)],
			its[len(insertTuples):]); err != nil {

		h.d.Writer().WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
