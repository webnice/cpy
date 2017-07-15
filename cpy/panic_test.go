package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestPanicRecovery(t *testing.T) {
	var err error
	type t1 struct {
		F1 int64
	}
	var src t1
	var dst int64

	if err = All(&dst, &src); err == nil {
		t.Fatal("Copy catch panic error")
	}
}
