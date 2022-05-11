// Package cpy
package cpy

// All Copy everything from one to another
func All(toObj interface{}, fromObj interface{}) error {
	return singleton.Copy(toObj, fromObj, nil, nil, nil)
}

// Select Copy only the selected fields.
// Use for struct only
func Select(toObj interface{}, fromObj interface{}, fields ...string) error {
	return singleton.Copy(toObj, fromObj, fields, nil, nil)
}

// Omit Copy everything from one to another, but skip listed fields.
// Use for struct only
func Omit(toObj interface{}, fromObj interface{}, fields ...string) error {
	return singleton.Copy(toObj, fromObj, nil, fields, nil)
}

// Filter Copy everything data which filtration, used for array, slice and map
func Filter(toObj interface{}, fromObj interface{}, filter FilterFn) error {
	return singleton.Copy(toObj, fromObj, nil, nil, filter)
}

// Errors

// ErrCopyToObjectUnaddressable Error: Copy to object is unaddressable
func ErrCopyToObjectUnaddressable() error { return singleton.ErrCopyToObjectUnaddressable() }

// ErrCopyFromObjectInvalid Error: Copy from object is invalid
func ErrCopyFromObjectInvalid() error { return singleton.ErrCopyFromObjectInvalid() }

// ErrTypeMapNotEqual Error: Type of map is not equal
func ErrTypeMapNotEqual() error { return singleton.ErrTypeMapNotEqual() }
