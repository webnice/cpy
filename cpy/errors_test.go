// Package cpy
package cpy

import "testing"

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
	var (
		err error
		src *One
	)

	dst := Two{}
	err = All(&dst, src)
	if err == nil || err != ErrCopyFromObjectInvalid() {
		t.Fatalf("Error check invalid value")
	}
}

func TestErrTypeMapNotEqual(t *testing.T) {
	type mt struct {
		I int64
		T string
	}
	var (
		err error
		m1  map[int64]mt
		m2  map[int64]*mt
	)

	m1 = make(map[int64]mt)
	m1[-1] = mt{T: "Minus one"}
	m1[100] = mt{I: 101, T: "One hundred"}
	err = All(&m2, &m1)
	if err == nil {
		t.Fatal("Copy map to map failed")
	}
	if err != ErrTypeMapNotEqual() {
		t.Fatalf("Incorrect error from copy map to map")
	}
}
