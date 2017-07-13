package cpy

//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/debug.v1"
import (
	"testing"
)

func TestAllEmbedded(t *testing.T) {
	var err error

	type (
		Destination struct {
			DestinationField1 int8
			DestinationField2 int64
		}
		Source struct {
			SourceField1 int8
			SourceField2 int64
			Destination
		}
	)

	dst := Destination{}
	src := Source{}
	src.DestinationField1 = 1
	src.DestinationField2 = 2
	src.SourceField1 = 3
	src.SourceField2 = 4

	err = All(&dst, &src)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if dst.DestinationField1 != 1 {
		t.Error("Embedded fields not copied")
	}
	if dst.DestinationField2 != 2 {
		t.Error("Embedded fields not copied")
	}

	debug.Dumper(dst)
}

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
