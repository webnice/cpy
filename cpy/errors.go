package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import "fmt"

var (
	errCopyToObjectUnaddressable = fmt.Errorf("Copy to object is unaddressable")
	errCopyFromObjectInvalid     = fmt.Errorf("Copy from object is invalid")
)

// ErrCopyToObjectUnaddressable Error: Copy to object is unaddressable
func (cpy *impl) ErrCopyToObjectUnaddressable() error { return errCopyToObjectUnaddressable }

// ErrCopyFromObjectInvalid Error: Copy from object is invalid
func (cpy *impl) ErrCopyFromObjectInvalid() error { return errCopyFromObjectInvalid }
