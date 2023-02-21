// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ory/keto/relation_tuples/v1alpha2/write_service.proto

package rts

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TransactRelationTuplesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TransactRelationTuplesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransactRelationTuplesRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// TransactRelationTuplesRequestMultiError, or nil if none found.
func (m *TransactRelationTuplesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TransactRelationTuplesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRelationTupleDeltas() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TransactRelationTuplesRequestValidationError{
						field:  fmt.Sprintf("RelationTupleDeltas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TransactRelationTuplesRequestValidationError{
						field:  fmt.Sprintf("RelationTupleDeltas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TransactRelationTuplesRequestValidationError{
					field:  fmt.Sprintf("RelationTupleDeltas[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TransactRelationTuplesRequestMultiError(errors)
	}

	return nil
}

// TransactRelationTuplesRequestMultiError is an error wrapping multiple
// validation errors returned by TransactRelationTuplesRequest.ValidateAll()
// if the designated constraints aren't met.
type TransactRelationTuplesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactRelationTuplesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactRelationTuplesRequestMultiError) AllErrors() []error { return m }

// TransactRelationTuplesRequestValidationError is the validation error
// returned by TransactRelationTuplesRequest.Validate if the designated
// constraints aren't met.
type TransactRelationTuplesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactRelationTuplesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactRelationTuplesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactRelationTuplesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactRelationTuplesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactRelationTuplesRequestValidationError) ErrorName() string {
	return "TransactRelationTuplesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TransactRelationTuplesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactRelationTuplesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactRelationTuplesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactRelationTuplesRequestValidationError{}

// Validate checks the field values on RelationTupleDelta with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RelationTupleDelta) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RelationTupleDelta with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RelationTupleDeltaMultiError, or nil if none found.
func (m *RelationTupleDelta) ValidateAll() error {
	return m.validate(true)
}

func (m *RelationTupleDelta) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _RelationTupleDelta_Action_InLookup[m.GetAction()]; !ok {
		err := RelationTupleDeltaValidationError{
			field:  "Action",
			reason: "value must be in list [1 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRelationTuple() == nil {
		err := RelationTupleDeltaValidationError{
			field:  "RelationTuple",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRelationTuple()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationTupleDeltaValidationError{
					field:  "RelationTuple",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationTupleDeltaValidationError{
					field:  "RelationTuple",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRelationTuple()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationTupleDeltaValidationError{
				field:  "RelationTuple",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RelationTupleDeltaMultiError(errors)
	}

	return nil
}

// RelationTupleDeltaMultiError is an error wrapping multiple validation errors
// returned by RelationTupleDelta.ValidateAll() if the designated constraints
// aren't met.
type RelationTupleDeltaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationTupleDeltaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationTupleDeltaMultiError) AllErrors() []error { return m }

// RelationTupleDeltaValidationError is the validation error returned by
// RelationTupleDelta.Validate if the designated constraints aren't met.
type RelationTupleDeltaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationTupleDeltaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationTupleDeltaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationTupleDeltaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationTupleDeltaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationTupleDeltaValidationError) ErrorName() string {
	return "RelationTupleDeltaValidationError"
}

// Error satisfies the builtin error interface
func (e RelationTupleDeltaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationTupleDelta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationTupleDeltaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationTupleDeltaValidationError{}

var _RelationTupleDelta_Action_InLookup = map[RelationTupleDelta_Action]struct{}{
	1: {},
	2: {},
}

// Validate checks the field values on TransactRelationTuplesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TransactRelationTuplesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransactRelationTuplesResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// TransactRelationTuplesResponseMultiError, or nil if none found.
func (m *TransactRelationTuplesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TransactRelationTuplesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return TransactRelationTuplesResponseMultiError(errors)
	}

	return nil
}

// TransactRelationTuplesResponseMultiError is an error wrapping multiple
// validation errors returned by TransactRelationTuplesResponse.ValidateAll()
// if the designated constraints aren't met.
type TransactRelationTuplesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactRelationTuplesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactRelationTuplesResponseMultiError) AllErrors() []error { return m }

// TransactRelationTuplesResponseValidationError is the validation error
// returned by TransactRelationTuplesResponse.Validate if the designated
// constraints aren't met.
type TransactRelationTuplesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactRelationTuplesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactRelationTuplesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactRelationTuplesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactRelationTuplesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactRelationTuplesResponseValidationError) ErrorName() string {
	return "TransactRelationTuplesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TransactRelationTuplesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactRelationTuplesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactRelationTuplesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactRelationTuplesResponseValidationError{}

// Validate checks the field values on CreateRelationTupleRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRelationTupleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRelationTupleRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRelationTupleRequestMultiError, or nil if none found.
func (m *CreateRelationTupleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRelationTupleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRelationTuple()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRelationTupleRequestValidationError{
					field:  "RelationTuple",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRelationTupleRequestValidationError{
					field:  "RelationTuple",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRelationTuple()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRelationTupleRequestValidationError{
				field:  "RelationTuple",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRelationTupleRequestMultiError(errors)
	}

	return nil
}

// CreateRelationTupleRequestMultiError is an error wrapping multiple
// validation errors returned by CreateRelationTupleRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateRelationTupleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRelationTupleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRelationTupleRequestMultiError) AllErrors() []error { return m }

// CreateRelationTupleRequestValidationError is the validation error returned
// by CreateRelationTupleRequest.Validate if the designated constraints aren't met.
type CreateRelationTupleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRelationTupleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRelationTupleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRelationTupleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRelationTupleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRelationTupleRequestValidationError) ErrorName() string {
	return "CreateRelationTupleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRelationTupleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRelationTupleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRelationTupleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRelationTupleRequestValidationError{}

// Validate checks the field values on CreateRelationTupleResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRelationTupleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRelationTupleResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRelationTupleResponseMultiError, or nil if none found.
func (m *CreateRelationTupleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRelationTupleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRelationTuple()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRelationTupleResponseValidationError{
					field:  "RelationTuple",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRelationTupleResponseValidationError{
					field:  "RelationTuple",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRelationTuple()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRelationTupleResponseValidationError{
				field:  "RelationTuple",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRelationTupleResponseMultiError(errors)
	}

	return nil
}

// CreateRelationTupleResponseMultiError is an error wrapping multiple
// validation errors returned by CreateRelationTupleResponse.ValidateAll() if
// the designated constraints aren't met.
type CreateRelationTupleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRelationTupleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRelationTupleResponseMultiError) AllErrors() []error { return m }

// CreateRelationTupleResponseValidationError is the validation error returned
// by CreateRelationTupleResponse.Validate if the designated constraints
// aren't met.
type CreateRelationTupleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRelationTupleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRelationTupleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRelationTupleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRelationTupleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRelationTupleResponseValidationError) ErrorName() string {
	return "CreateRelationTupleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRelationTupleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRelationTupleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRelationTupleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRelationTupleResponseValidationError{}

// Validate checks the field values on DeleteRelationTuplesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRelationTuplesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRelationTuplesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRelationTuplesRequestMultiError, or nil if none found.
func (m *DeleteRelationTuplesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRelationTuplesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuery()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteRelationTuplesRequestValidationError{
					field:  "Query",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteRelationTuplesRequestValidationError{
					field:  "Query",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteRelationTuplesRequestValidationError{
				field:  "Query",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRelationQuery()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteRelationTuplesRequestValidationError{
					field:  "RelationQuery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteRelationTuplesRequestValidationError{
					field:  "RelationQuery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRelationQuery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteRelationTuplesRequestValidationError{
				field:  "RelationQuery",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Namespace

	// no validation rules for Object

	// no validation rules for Relation

	switch v := m.RestApiSubject.(type) {
	case *DeleteRelationTuplesRequest_SubjectId:
		if v == nil {
			err := DeleteRelationTuplesRequestValidationError{
				field:  "RestApiSubject",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for SubjectId
	case *DeleteRelationTuplesRequest_SubjectSet:
		if v == nil {
			err := DeleteRelationTuplesRequestValidationError{
				field:  "RestApiSubject",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSubjectSet()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeleteRelationTuplesRequestValidationError{
						field:  "SubjectSet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeleteRelationTuplesRequestValidationError{
						field:  "SubjectSet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSubjectSet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeleteRelationTuplesRequestValidationError{
					field:  "SubjectSet",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DeleteRelationTuplesRequestMultiError(errors)
	}

	return nil
}

// DeleteRelationTuplesRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteRelationTuplesRequest.ValidateAll() if
// the designated constraints aren't met.
type DeleteRelationTuplesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRelationTuplesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRelationTuplesRequestMultiError) AllErrors() []error { return m }

// DeleteRelationTuplesRequestValidationError is the validation error returned
// by DeleteRelationTuplesRequest.Validate if the designated constraints
// aren't met.
type DeleteRelationTuplesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRelationTuplesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRelationTuplesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRelationTuplesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRelationTuplesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRelationTuplesRequestValidationError) ErrorName() string {
	return "DeleteRelationTuplesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRelationTuplesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRelationTuplesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRelationTuplesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRelationTuplesRequestValidationError{}

// Validate checks the field values on DeleteRelationTuplesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRelationTuplesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRelationTuplesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRelationTuplesResponseMultiError, or nil if none found.
func (m *DeleteRelationTuplesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRelationTuplesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRelationTuplesResponseMultiError(errors)
	}

	return nil
}

// DeleteRelationTuplesResponseMultiError is an error wrapping multiple
// validation errors returned by DeleteRelationTuplesResponse.ValidateAll() if
// the designated constraints aren't met.
type DeleteRelationTuplesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRelationTuplesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRelationTuplesResponseMultiError) AllErrors() []error { return m }

// DeleteRelationTuplesResponseValidationError is the validation error returned
// by DeleteRelationTuplesResponse.Validate if the designated constraints
// aren't met.
type DeleteRelationTuplesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRelationTuplesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRelationTuplesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRelationTuplesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRelationTuplesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRelationTuplesResponseValidationError) ErrorName() string {
	return "DeleteRelationTuplesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRelationTuplesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRelationTuplesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRelationTuplesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRelationTuplesResponseValidationError{}

// Validate checks the field values on CreateRelationTupleRequest_Relationship
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CreateRelationTupleRequest_Relationship) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CreateRelationTupleRequest_Relationship with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// CreateRelationTupleRequest_RelationshipMultiError, or nil if none found.
func (m *CreateRelationTupleRequest_Relationship) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRelationTupleRequest_Relationship) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Object

	// no validation rules for Relation

	switch v := m.Subject.(type) {
	case *CreateRelationTupleRequest_Relationship_SubjectId:
		if v == nil {
			err := CreateRelationTupleRequest_RelationshipValidationError{
				field:  "Subject",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for SubjectId
	case *CreateRelationTupleRequest_Relationship_SubjectSet:
		if v == nil {
			err := CreateRelationTupleRequest_RelationshipValidationError{
				field:  "Subject",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSubjectSet()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateRelationTupleRequest_RelationshipValidationError{
						field:  "SubjectSet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateRelationTupleRequest_RelationshipValidationError{
						field:  "SubjectSet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSubjectSet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateRelationTupleRequest_RelationshipValidationError{
					field:  "SubjectSet",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return CreateRelationTupleRequest_RelationshipMultiError(errors)
	}

	return nil
}

// CreateRelationTupleRequest_RelationshipMultiError is an error wrapping
// multiple validation errors returned by
// CreateRelationTupleRequest_Relationship.ValidateAll() if the designated
// constraints aren't met.
type CreateRelationTupleRequest_RelationshipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRelationTupleRequest_RelationshipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRelationTupleRequest_RelationshipMultiError) AllErrors() []error { return m }

// CreateRelationTupleRequest_RelationshipValidationError is the validation
// error returned by CreateRelationTupleRequest_Relationship.Validate if the
// designated constraints aren't met.
type CreateRelationTupleRequest_RelationshipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRelationTupleRequest_RelationshipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRelationTupleRequest_RelationshipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRelationTupleRequest_RelationshipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRelationTupleRequest_RelationshipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRelationTupleRequest_RelationshipValidationError) ErrorName() string {
	return "CreateRelationTupleRequest_RelationshipValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRelationTupleRequest_RelationshipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRelationTupleRequest_Relationship.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRelationTupleRequest_RelationshipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRelationTupleRequest_RelationshipValidationError{}

// Validate checks the field values on DeleteRelationTuplesRequest_Query with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *DeleteRelationTuplesRequest_Query) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRelationTuplesRequest_Query
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DeleteRelationTuplesRequest_QueryMultiError, or nil if none found.
func (m *DeleteRelationTuplesRequest_Query) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRelationTuplesRequest_Query) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Object

	// no validation rules for Relation

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteRelationTuplesRequest_QueryValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteRelationTuplesRequest_QueryValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteRelationTuplesRequest_QueryValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteRelationTuplesRequest_QueryMultiError(errors)
	}

	return nil
}

// DeleteRelationTuplesRequest_QueryMultiError is an error wrapping multiple
// validation errors returned by
// DeleteRelationTuplesRequest_Query.ValidateAll() if the designated
// constraints aren't met.
type DeleteRelationTuplesRequest_QueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRelationTuplesRequest_QueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRelationTuplesRequest_QueryMultiError) AllErrors() []error { return m }

// DeleteRelationTuplesRequest_QueryValidationError is the validation error
// returned by DeleteRelationTuplesRequest_Query.Validate if the designated
// constraints aren't met.
type DeleteRelationTuplesRequest_QueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRelationTuplesRequest_QueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRelationTuplesRequest_QueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRelationTuplesRequest_QueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRelationTuplesRequest_QueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRelationTuplesRequest_QueryValidationError) ErrorName() string {
	return "DeleteRelationTuplesRequest_QueryValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRelationTuplesRequest_QueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRelationTuplesRequest_Query.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRelationTuplesRequest_QueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRelationTuplesRequest_QueryValidationError{}
