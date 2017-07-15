package cpy

//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/debug.v1"
import (
	"testing"
	"time"
)

func init() {
	debug.Nop()
}

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
		t.Fatalf("Embedded fields not copied")
	}
	if dst.DestinationField2 != 2 {
		t.Fatalf("Embedded fields not copied")
	}
}

func TestMapAll(t *testing.T) {
	var err error
	type mt struct {
		I int64
		T string
	}
	var m1 map[int64]*mt
	var m2 map[int64]*mt
	var m3 map[string]mt
	var m4 map[string]mt

	m1 = make(map[int64]*mt)
	m1[-1] = &mt{T: "Minus one"}
	m1[100] = &mt{I: 101, T: "One hundred"}
	err = All(&m2, &m1)
	if err != nil {
		t.Fatalf("Copy map to map failed: %s", err.Error())
	}
	if v, ok := m2[-1]; !ok || v.T != "Minus one" {
		t.Fatalf("Copy map to map failed")
	}
	if v, ok := m2[100]; !ok || v.T != "One hundred" || v.I != 101 {
		t.Fatalf("Copy map to map failed")
	}

	m3 = make(map[string]mt)
	m3["-1"] = mt{T: "Minus one"}
	m3["100"] = mt{I: 101, T: "One hundred"}
	err = All(&m4, &m3)
	if err != nil {
		t.Fatalf("Copy map to map failed: %s", err.Error())
	}
	if v, ok := m4["-1"]; !ok || v.T != "Minus one" {
		t.Fatalf("Copy map to map failed")
	}
	if v, ok := m4["100"]; !ok || v.T != "One hundred" || v.I != 101 {
		t.Fatalf("Copy map to map failed")
	}
}

func TestAllConverting(t *testing.T) {
	var err error
	var src *One
	var dst *Converting
	var tm time.Time

	src = createOne()
	dst = new(Converting)
	err = All(dst, src)
	if err != nil {
		t.Fatalf("Copy All failed: %s", err.Error())
	}
	if dst.NewID != 1 {
		t.Fatal("Copy All failed")
	}
	if dst.Int64 != -1234567 {
		t.Fatal("Copy All failed")
	}
	if dst.Cat != "myau" {
		t.Fatal("Copy All failed")
	}
	tm, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-07-15 02:08:46.691821235 +0000 UTC")
	if !dst.Time.Time.Equal(tm) {
		t.Fatal("Copy All failed")
	}
}

func TestAllSlice(t *testing.T) {
	var err error
	var src1 []*One
	var src2 []One
	var dst1 []Two
	var dst2 []*Two

	tmp := createOne()
	src1 = []*One{tmp, tmp, tmp}
	src2 = []One{*tmp, *tmp, *tmp}
	if err = All(&dst1, &src1); err != nil {
		t.Fatalf("Copy slice failed: %s", err.Error())
	}
	if err = All(&dst2, &src2); err != nil {
		t.Fatalf("Copy slice failed: %s", err.Error())
	}
	if len(dst1) != len(src1) || len(dst2) != len(src2) {
		t.Fatal("Copy All failed")
	}
}

func TestAllStructToSlice(t *testing.T) {
	var err error
	var src *One
	var dst []Two

	src = createOne()
	if err = All(&dst, &src); err != nil {
		t.Fatalf("Copy slice failed: %s", err.Error())
	}
	if len(dst) != 1 {
		t.Fatal("Copy All failed")
	}
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

	//debug.Dumper(src, dst)
}
