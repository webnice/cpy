// Package cpy
package cpy

import "testing"

func TestPanicRecovery(t *testing.T) {
	type t1 struct {
		F1 int64
	}
	var (
		err error
		src t1
		dst int64
	)

	if err = All(&dst, &src); err == nil {
		t.Fatal("Copy catch panic error")
	}
}
