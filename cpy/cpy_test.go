package cpy

//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/debug.v1"
import (
	"testing"
)

func TestAll(t *testing.T) {
	var err error
	var src *One
	var dst *Two

	src = createOne()
	dst = new(Two)
	err = All(dst, src)
	if err != nil {
		t.Fatalf("Copy All failed: %s", err.Error())
	}

	debug.Dumper(src, dst)
}
