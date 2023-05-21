// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/storage/v1/storage.proto

package v1

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

// Validate checks the field values on StorageWriteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageWriteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageWriteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageWriteRequestMultiError, or nil if none found.
func (m *StorageWriteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageWriteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BucketName

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := StorageWriteRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Body

	if len(errors) > 0 {
		return StorageWriteRequestMultiError(errors)
	}

	return nil
}

// StorageWriteRequestMultiError is an error wrapping multiple validation
// errors returned by StorageWriteRequest.ValidateAll() if the designated
// constraints aren't met.
type StorageWriteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageWriteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageWriteRequestMultiError) AllErrors() []error { return m }

// StorageWriteRequestValidationError is the validation error returned by
// StorageWriteRequest.Validate if the designated constraints aren't met.
type StorageWriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageWriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageWriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageWriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageWriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageWriteRequestValidationError) ErrorName() string {
	return "StorageWriteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StorageWriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageWriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageWriteRequestValidationError{}

// Validate checks the field values on StorageWriteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageWriteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageWriteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageWriteResponseMultiError, or nil if none found.
func (m *StorageWriteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageWriteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StorageWriteResponseMultiError(errors)
	}

	return nil
}

// StorageWriteResponseMultiError is an error wrapping multiple validation
// errors returned by StorageWriteResponse.ValidateAll() if the designated
// constraints aren't met.
type StorageWriteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageWriteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageWriteResponseMultiError) AllErrors() []error { return m }

// StorageWriteResponseValidationError is the validation error returned by
// StorageWriteResponse.Validate if the designated constraints aren't met.
type StorageWriteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageWriteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageWriteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageWriteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageWriteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageWriteResponseValidationError) ErrorName() string {
	return "StorageWriteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StorageWriteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageWriteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageWriteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageWriteResponseValidationError{}

// Validate checks the field values on StorageReadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageReadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageReadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageReadRequestMultiError, or nil if none found.
func (m *StorageReadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageReadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BucketName

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := StorageReadRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StorageReadRequestMultiError(errors)
	}

	return nil
}

// StorageReadRequestMultiError is an error wrapping multiple validation errors
// returned by StorageReadRequest.ValidateAll() if the designated constraints
// aren't met.
type StorageReadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageReadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageReadRequestMultiError) AllErrors() []error { return m }

// StorageReadRequestValidationError is the validation error returned by
// StorageReadRequest.Validate if the designated constraints aren't met.
type StorageReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageReadRequestValidationError) ErrorName() string {
	return "StorageReadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StorageReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageReadRequestValidationError{}

// Validate checks the field values on StorageReadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageReadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageReadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageReadResponseMultiError, or nil if none found.
func (m *StorageReadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageReadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Body

	if len(errors) > 0 {
		return StorageReadResponseMultiError(errors)
	}

	return nil
}

// StorageReadResponseMultiError is an error wrapping multiple validation
// errors returned by StorageReadResponse.ValidateAll() if the designated
// constraints aren't met.
type StorageReadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageReadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageReadResponseMultiError) AllErrors() []error { return m }

// StorageReadResponseValidationError is the validation error returned by
// StorageReadResponse.Validate if the designated constraints aren't met.
type StorageReadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageReadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageReadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageReadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageReadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageReadResponseValidationError) ErrorName() string {
	return "StorageReadResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StorageReadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageReadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageReadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageReadResponseValidationError{}

// Validate checks the field values on StorageDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageDeleteRequestMultiError, or nil if none found.
func (m *StorageDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BucketName

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := StorageDeleteRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StorageDeleteRequestMultiError(errors)
	}

	return nil
}

// StorageDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by StorageDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type StorageDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageDeleteRequestMultiError) AllErrors() []error { return m }

// StorageDeleteRequestValidationError is the validation error returned by
// StorageDeleteRequest.Validate if the designated constraints aren't met.
type StorageDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageDeleteRequestValidationError) ErrorName() string {
	return "StorageDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StorageDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageDeleteRequestValidationError{}

// Validate checks the field values on StorageDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageDeleteResponseMultiError, or nil if none found.
func (m *StorageDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StorageDeleteResponseMultiError(errors)
	}

	return nil
}

// StorageDeleteResponseMultiError is an error wrapping multiple validation
// errors returned by StorageDeleteResponse.ValidateAll() if the designated
// constraints aren't met.
type StorageDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageDeleteResponseMultiError) AllErrors() []error { return m }

// StorageDeleteResponseValidationError is the validation error returned by
// StorageDeleteResponse.Validate if the designated constraints aren't met.
type StorageDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageDeleteResponseValidationError) ErrorName() string {
	return "StorageDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StorageDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageDeleteResponseValidationError{}

// Validate checks the field values on StoragePreSignUrlRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StoragePreSignUrlRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StoragePreSignUrlRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StoragePreSignUrlRequestMultiError, or nil if none found.
func (m *StoragePreSignUrlRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StoragePreSignUrlRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BucketName

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := StoragePreSignUrlRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Operation

	// no validation rules for Expiry

	if len(errors) > 0 {
		return StoragePreSignUrlRequestMultiError(errors)
	}

	return nil
}

// StoragePreSignUrlRequestMultiError is an error wrapping multiple validation
// errors returned by StoragePreSignUrlRequest.ValidateAll() if the designated
// constraints aren't met.
type StoragePreSignUrlRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StoragePreSignUrlRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StoragePreSignUrlRequestMultiError) AllErrors() []error { return m }

// StoragePreSignUrlRequestValidationError is the validation error returned by
// StoragePreSignUrlRequest.Validate if the designated constraints aren't met.
type StoragePreSignUrlRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StoragePreSignUrlRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StoragePreSignUrlRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StoragePreSignUrlRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StoragePreSignUrlRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StoragePreSignUrlRequestValidationError) ErrorName() string {
	return "StoragePreSignUrlRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StoragePreSignUrlRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStoragePreSignUrlRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StoragePreSignUrlRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StoragePreSignUrlRequestValidationError{}

// Validate checks the field values on StoragePreSignUrlResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StoragePreSignUrlResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StoragePreSignUrlResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StoragePreSignUrlResponseMultiError, or nil if none found.
func (m *StoragePreSignUrlResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StoragePreSignUrlResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return StoragePreSignUrlResponseMultiError(errors)
	}

	return nil
}

// StoragePreSignUrlResponseMultiError is an error wrapping multiple validation
// errors returned by StoragePreSignUrlResponse.ValidateAll() if the
// designated constraints aren't met.
type StoragePreSignUrlResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StoragePreSignUrlResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StoragePreSignUrlResponseMultiError) AllErrors() []error { return m }

// StoragePreSignUrlResponseValidationError is the validation error returned by
// StoragePreSignUrlResponse.Validate if the designated constraints aren't met.
type StoragePreSignUrlResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StoragePreSignUrlResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StoragePreSignUrlResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StoragePreSignUrlResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StoragePreSignUrlResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StoragePreSignUrlResponseValidationError) ErrorName() string {
	return "StoragePreSignUrlResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StoragePreSignUrlResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStoragePreSignUrlResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StoragePreSignUrlResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StoragePreSignUrlResponseValidationError{}

// Validate checks the field values on StorageListFilesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageListFilesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageListFilesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageListFilesRequestMultiError, or nil if none found.
func (m *StorageListFilesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageListFilesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BucketName

	// no validation rules for Prefix

	if len(errors) > 0 {
		return StorageListFilesRequestMultiError(errors)
	}

	return nil
}

// StorageListFilesRequestMultiError is an error wrapping multiple validation
// errors returned by StorageListFilesRequest.ValidateAll() if the designated
// constraints aren't met.
type StorageListFilesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageListFilesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageListFilesRequestMultiError) AllErrors() []error { return m }

// StorageListFilesRequestValidationError is the validation error returned by
// StorageListFilesRequest.Validate if the designated constraints aren't met.
type StorageListFilesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageListFilesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageListFilesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageListFilesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageListFilesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageListFilesRequestValidationError) ErrorName() string {
	return "StorageListFilesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StorageListFilesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageListFilesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageListFilesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageListFilesRequestValidationError{}

// Validate checks the field values on File with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *File) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on File with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FileMultiError, or nil if none found.
func (m *File) ValidateAll() error {
	return m.validate(true)
}

func (m *File) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return FileMultiError(errors)
	}

	return nil
}

// FileMultiError is an error wrapping multiple validation errors returned by
// File.ValidateAll() if the designated constraints aren't met.
type FileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileMultiError) AllErrors() []error { return m }

// FileValidationError is the validation error returned by File.Validate if the
// designated constraints aren't met.
type FileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileValidationError) ErrorName() string { return "FileValidationError" }

// Error satisfies the builtin error interface
func (e FileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileValidationError{}

// Validate checks the field values on StorageListFilesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StorageListFilesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageListFilesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageListFilesResponseMultiError, or nil if none found.
func (m *StorageListFilesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageListFilesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetFiles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StorageListFilesResponseValidationError{
						field:  fmt.Sprintf("Files[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StorageListFilesResponseValidationError{
						field:  fmt.Sprintf("Files[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StorageListFilesResponseValidationError{
					field:  fmt.Sprintf("Files[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StorageListFilesResponseMultiError(errors)
	}

	return nil
}

// StorageListFilesResponseMultiError is an error wrapping multiple validation
// errors returned by StorageListFilesResponse.ValidateAll() if the designated
// constraints aren't met.
type StorageListFilesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageListFilesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageListFilesResponseMultiError) AllErrors() []error { return m }

// StorageListFilesResponseValidationError is the validation error returned by
// StorageListFilesResponse.Validate if the designated constraints aren't met.
type StorageListFilesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageListFilesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageListFilesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageListFilesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageListFilesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageListFilesResponseValidationError) ErrorName() string {
	return "StorageListFilesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StorageListFilesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageListFilesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageListFilesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageListFilesResponseValidationError{}
