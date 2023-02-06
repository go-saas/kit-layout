// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: private/conf/conf.proto

package conf

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

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Bootstrap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bootstrap with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BootstrapMultiError, or nil
// if none found.
func (m *Bootstrap) ValidateAll() error {
	return m.validate(true)
}

func (m *Bootstrap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSecurity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Security",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Security",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSecurity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Security",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetServices()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Services",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Services",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServices()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Services",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLogging()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Logging",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Logging",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLogging()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Logging",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTracing()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Tracing",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Tracing",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTracing()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Tracing",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetApp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDev()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Dev",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Dev",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDev()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Dev",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BootstrapMultiError(errors)
	}

	return nil
}

// BootstrapMultiError is an error wrapping multiple validation errors returned
// by Bootstrap.ValidateAll() if the designated constraints aren't met.
type BootstrapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BootstrapMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BootstrapMultiError) AllErrors() []error { return m }

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}
