package cpy

import (
	"testing"
)

func TestErrCopyToObjectUnaddressable(t *testing.T) {
	var err error

	src := createOne()
	dst := Two{}
	err = All(dst, src)
	if err == nil || err != ErrCopyToObjectUnaddressable() {
		t.Fatalf("Error check unaddressable value")
	}
}

func TestErrCopyFromObjectInvalid(t *testing.T) {
	var err error
	var src *One

	dst := Two{}
	err = All(&dst, src)
	if err == nil || err != ErrCopyFromObjectInvalid() {
		t.Fatalf("Error check invalid value")
	}
}
