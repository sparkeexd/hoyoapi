package models

import (
	"bytes"
	"encoding/json"

	"github.com/sparkeexd/hoyoapi/internal/errors"
)

// Code taken from https://github.com/oapi-codegen/nullable/blob/v1.0.0/nullable.go
//
// Nullable is a generic type, which implements a property that can be one of three states:
// - Property is not set in the request
// - Property is explicitly set to `null` in the request
// - Property is explicitly set to a valid value in the request
//
// Nullable is intended to be used with JSON marshalling and unmarshalling.
//
// Internal implementation details:
// - map[true]T means a value is provided.
// - map[false]T means an explicit `null` is provided.
// - nil or zero map means the property is not provided.
//
// If the property is expected to be optional, add the `omitempty` JSON tags. Do NOT use `*Nullable`!
//
// Adapted from https://github.com/golang/go/issues/64515#issuecomment-1841057182
type Nullable[T any] map[bool]T

// Retrieves the property's underlying value, if present, and returns an error if the value is not present.
func (t Nullable[T]) Get() (T, error) {
	var empty T

	if t.IsNull() {
		return empty, errors.NewError(errors.NULLABLE_VALUE_ERROR, "Value is null")
	}

	if !t.IsDefined() {
		return empty, errors.NewError(errors.NULLABLE_VALUE_ERROR, "Value is undefined")
	}

	return t[true], nil
}

// Sets the property's underlying value.
func (t *Nullable[T]) Set(value T) {
	*t = map[bool]T{true: value}
}

// Check whether the property is present, and has the value of `null`.
func (t Nullable[T]) IsNull() bool {
	_, foundNull := t[false]
	return foundNull
}

// The property is present, and set its underlying value as `null`.
func (t *Nullable[T]) SetNull() {
	var empty T
	*t = map[bool]T{false: empty}
}

// Check whether the property is present.
func (t Nullable[T]) IsDefined() bool {
	return len(t) != 0
}

// Set the property to be undefined.
func (t *Nullable[T]) SetUndefined() {
	*t = map[bool]T{}
}

// Marshal JSON.
func (t Nullable[T]) MarshalJSON() ([]byte, error) {
	// If property is defined and value is `null`, marshal it.
	if t.IsNull() {
		return []byte("null"), nil
	}

	// If property is undefined and `omitempty` is set on the property's tags, `json.Marshal` will omit this property.
	// Otherwise, we have a value, so marshal it.
	return json.Marshal(t[true])
}

// Unmarshal JSON.
// If the property is undefined, UnmarshalJSON will not be called.
func (t *Nullable[T]) UnmarshalJSON(data []byte) error {

	// If property is specified, and `null`
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}

	// Otherwise, we have an actual value, so parse it.
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.Set(v)
	return nil
}
