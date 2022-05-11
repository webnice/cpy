// Package cpy
package cpy

import "fmt"

var (
	errCopyToObjectUnaddressable = fmt.Errorf("copy to object is unaddressable")
	errCopyFromObjectInvalid     = fmt.Errorf("copy from object is invalid")
	errTypeMapNotEqual           = fmt.Errorf("type of map is not equal")
)

// ErrCopyToObjectUnaddressable Error: Copy to object is unaddressable
func (cpy *impl) ErrCopyToObjectUnaddressable() error { return errCopyToObjectUnaddressable }

// ErrCopyFromObjectInvalid Error: Copy from object is invalid
func (cpy *impl) ErrCopyFromObjectInvalid() error { return errCopyFromObjectInvalid }

// ErrTypeMapNotEqual Error: Type of map is not equal
func (cpy *impl) ErrTypeMapNotEqual() error { return errTypeMapNotEqual }
