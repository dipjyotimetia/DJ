// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/discovery/v2/rtds.proto

package discoveryv2

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

// Validate checks the field values on RtdsDummy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RtdsDummy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RtdsDummy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RtdsDummyMultiError, or nil
// if none found.
func (m *RtdsDummy) ValidateAll() error {
	return m.validate(true)
}

func (m *RtdsDummy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RtdsDummyMultiError(errors)
	}

	return nil
}

// RtdsDummyMultiError is an error wrapping multiple validation errors returned
// by RtdsDummy.ValidateAll() if the designated constraints aren't met.
type RtdsDummyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RtdsDummyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RtdsDummyMultiError) AllErrors() []error { return m }

// RtdsDummyValidationError is the validation error returned by
// RtdsDummy.Validate if the designated constraints aren't met.
type RtdsDummyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RtdsDummyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RtdsDummyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RtdsDummyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RtdsDummyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RtdsDummyValidationError) ErrorName() string { return "RtdsDummyValidationError" }

// Error satisfies the builtin error interface
func (e RtdsDummyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRtdsDummy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RtdsDummyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RtdsDummyValidationError{}

// Validate checks the field values on Runtime with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Runtime) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Runtime with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RuntimeMultiError, or nil if none found.
func (m *Runtime) ValidateAll() error {
	return m.validate(true)
}

func (m *Runtime) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) < 1 {
		err := RuntimeValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetLayer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RuntimeValidationError{
					field:  "Layer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RuntimeValidationError{
					field:  "Layer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLayer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuntimeValidationError{
				field:  "Layer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RuntimeMultiError(errors)
	}

	return nil
}

// RuntimeMultiError is an error wrapping multiple validation errors returned
// by Runtime.ValidateAll() if the designated constraints aren't met.
type RuntimeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RuntimeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RuntimeMultiError) AllErrors() []error { return m }

// RuntimeValidationError is the validation error returned by Runtime.Validate
// if the designated constraints aren't met.
type RuntimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeValidationError) ErrorName() string { return "RuntimeValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeValidationError{}
