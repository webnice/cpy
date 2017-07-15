package cpy

//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/debug.v1"
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

func TestErrTypeMapNotEqual(t *testing.T) {
	var err error
	type mt struct {
		I int64
		T string
	}
	var m1 map[int64]mt
	var m2 map[int64]*mt

	m1 = make(map[int64]mt)
	m1[-1] = mt{T: "Minus one"}
	m1[100] = mt{I: 101, T: "One hundred"}
	if err = All(&m2, &m1); err == nil {

		debug.Dumper(err, m1, m2)

		t.Fatal("Copy map to map failed")
	}
}
