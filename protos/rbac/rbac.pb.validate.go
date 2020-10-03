// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rbac.proto

package rbac

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _rbac_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Request) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) > 36 {
		return RequestValidationError{
			field:  "Id",
			reason: "value length must be at most 36 runes",
		}
	}

	return nil
}

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on LinkRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LinkRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId1()) > 36 {
		return LinkRequestValidationError{
			field:  "Id1",
			reason: "value length must be at most 36 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId2()) > 36 {
		return LinkRequestValidationError{
			field:  "Id2",
			reason: "value length must be at most 36 runes",
		}
	}

	return nil
}

// LinkRequestValidationError is the validation error returned by
// LinkRequest.Validate if the designated constraints aren't met.
type LinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkRequestValidationError) ErrorName() string { return "LinkRequestValidationError" }

// Error satisfies the builtin error interface
func (e LinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkRequestValidationError{}

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Msg

	return nil
}

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) > 36 {
		return UserValidationError{
			field:  "Id",
			reason: "value length must be at most 36 runes",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 10 {
		return UserValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 10 runes, inclusive",
		}
	}

	if val := m.GetAge(); val <= 0 || val > 180 {
		return UserValidationError{
			field:  "Age",
			reason: "value must be inside range (0, 180]",
		}
	}

	if _, ok := _User_Gender_InLookup[m.GetGender()]; !ok {
		return UserValidationError{
			field:  "Gender",
			reason: "value must be in list [0 1 2]",
		}
	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

var _User_Gender_InLookup = map[int32]struct{}{
	0: {},
	1: {},
	2: {},
}

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Role) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) > 36 {
		return RoleValidationError{
			field:  "Id",
			reason: "value length must be at most 36 runes",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 10 {
		return RoleValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 10 runes, inclusive",
		}
	}

	return nil
}

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on Roles with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Roles) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RolesValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RolesValidationError is the validation error returned by Roles.Validate if
// the designated constraints aren't met.
type RolesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RolesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RolesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RolesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RolesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RolesValidationError) ErrorName() string { return "RolesValidationError" }

// Error satisfies the builtin error interface
func (e RolesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoles.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RolesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RolesValidationError{}

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resource) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) > 36 {
		return ResourceValidationError{
			field:  "Id",
			reason: "value length must be at most 36 runes",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 10 {
		return ResourceValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 10 runes, inclusive",
		}
	}

	return nil
}

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on Resources with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resources) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourcesValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResourcesValidationError is the validation error returned by
// Resources.Validate if the designated constraints aren't met.
type ResourcesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourcesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourcesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourcesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourcesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourcesValidationError) ErrorName() string { return "ResourcesValidationError" }

// Error satisfies the builtin error interface
func (e ResourcesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResources.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourcesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourcesValidationError{}