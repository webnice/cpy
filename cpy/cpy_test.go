// Package cpy
package cpy

import (
	"bytes"
	"testing"
	"time"
)

func TestAllEmbedded(t *testing.T) {
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
	var (
		err error
		dst Destination
		src Source
	)

	src.DestinationField1 = 1
	src.DestinationField2 = 2
	src.SourceField1 = 3
	src.SourceField2 = 4
	if err = All(&dst, &src); err != nil {
		t.Fatalf("error: %s", err)
	}
	if dst.DestinationField1 != 1 {
		t.Fatalf("embedded fields not copied")
	}
	if dst.DestinationField2 != 2 {
		t.Fatalf("embedded fields not copied")
	}
}

func TestMapAll(t *testing.T) {
	type mt struct {
		I int64
		T string
	}
	var (
		err error
		m1  map[int64]*mt
		m2  map[int64]*mt
		m3  map[string]mt
		m4  map[string]mt
		v   *mt
		ok  bool
	)

	m1 = make(map[int64]*mt)
	m1[-1] = &mt{T: "minus one"}
	m1[100] = &mt{I: 101, T: "one hundred"}
	if err = All(&m2, &m1); err != nil {
		t.Fatalf("copy map to map failed: %s", err)
	}
	if v, ok = m2[-1]; !ok || v.T != "minus one" {
		t.Fatalf("copy map to map failed")
	}
	if v, ok = m2[100]; !ok || v.T != "one hundred" || v.I != 101 {
		t.Fatalf("copy map to map failed")
	}
	m3 = make(map[string]mt)
	m3["-1"] = mt{T: "minus one"}
	m3["100"] = mt{I: 101, T: "one hundred"}
	if err = All(&m4, &m3); err != nil {
		t.Fatalf("copy map to map failed: %s", err)
	}
	if v, ok := m4["-1"]; !ok || v.T != "Minus one" {
		t.Fatalf("copy map to map failed")
	}
	if v, ok := m4["100"]; !ok || v.T != "one hundred" || v.I != 101 {
		t.Fatalf("copy map to map failed")
	}
}

func TestAllConverting(t *testing.T) {
	var (
		err error
		src *One
		dst *Converting
		tm  time.Time
	)

	src = createOne()
	dst = new(Converting)
	if err = All(dst, src); err != nil {
		t.Fatalf("copy all() failed: %s", err)
	}
	if dst.NewID != 1 {
		t.Fatal("copy all() failed")
	}
	if dst.Int64 != -1234567 {
		t.Fatal("copy all() failed")
	}
	if dst.Cat != "my-au" {
		t.Fatal("copy all() failed")
	}
	tm, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-07-15 02:08:46.691821235 +0000 UTC")
	if !dst.Time.Time.Equal(tm) {
		t.Fatal("copy all() failed")
	}
}

func TestAllSlice(t *testing.T) {
	var (
		err  error
		tmp  *One
		src1 []*One
		src2 []One
		dst1 []Two
		dst2 []*Two
	)

	tmp = createOne()
	src1 = []*One{tmp, tmp, tmp}
	src2 = []One{*tmp, *tmp, *tmp}
	if err = All(&dst1, &src1); err != nil {
		t.Fatalf("copy slice failed: %s", err)
	}
	if err = All(&dst2, &src2); err != nil {
		t.Fatalf("copy slice failed: %s", err)
	}
	if len(dst1) != len(src1) || len(dst2) != len(src2) {
		t.Fatal("copy all() failed")
	}
}

func TestAllStructToSlice(t *testing.T) {
	var (
		err error
		src *One
		dst []Two
	)

	src = createOne()
	if err = All(&dst, &src); err != nil {
		t.Fatalf("copy slice failed: %s", err)
	}
	if len(dst) != 1 {
		t.Fatal("copy all() failed")
	}
}

func TestAll(t *testing.T) {
	var (
		err error
		src *One
		dst *Two
	)

	src = createOne()
	dst = new(Two)
	err = All(dst, src)
	if err != nil {
		t.Fatalf("copy all() failed: %s", err)
	}
	if *dst.NewID != 1 {
		t.Fatal("copy all() failed")
	}
	if *dst.Name != "hello from One.Name" {
		t.Fatal("copy all() failed")
	}
	if !bytes.Equal(dst.Des, []byte("One.Description")) {
		t.Fatal("copy all() failed")
	}
	if dst.Complex != "One.Description, name: hello from One.Name" {
		t.Fatal("copy all() failed")
	}
	if !dst.Disabled {
		t.Fatal("copy all() failed")
	}
}

func TestSelect(t *testing.T) {
	var (
		err error
		src *One
		dst *Two
	)

	src = createOne()
	dst = new(Two)
	err = Select(dst, src, "ID", "Des")
	if err != nil {
		t.Fatalf("copy select() failed: %s", err)
	}
	if *dst.NewID != 1 {
		t.Fatal("copy select() failed")
	}
	if !bytes.Equal(dst.Des, []byte("One.Description")) {
		t.Fatal("copy select() failed")
	}
	if dst.Name != nil {
		t.Fatal("copy select() failed")
	}
	if dst.Complex == "One.Description, name: Hello from One.Name" {
		t.Fatal("copy select() failed")
	}
	if dst.Disabled {
		t.Fatal("copy select() failed")
	}
}

func TestOmit(t *testing.T) {
	var (
		err error
		src *One
		dst *Two
	)

	src = createOne()
	dst = new(Two)
	err = Omit(dst, src, "ID")
	if err != nil {
		t.Fatalf("copy omit() failed: %s", err)
	}
	if *dst.Name != "Hello from One.Name" {
		t.Fatal("copy omit() failed")
	}
	if !bytes.Equal(dst.Des, []byte("One.Description")) {
		t.Fatal("copy omit() failed")
	}
	if dst.Complex != "One.Description, name: Hello from One.Name" {
		t.Fatal("copy omit() failed")
	}
	if !dst.Disabled {
		t.Fatal("copy omit() failed")
	}

	if dst.NewID != nil {
		t.Fatal("copy omit() failed")
	}
}

func TestFilterSlice(t *testing.T) {
	var (
		err error
		src []*TFilter
		dst []*TFilter
		sum int64
	)

	src = createSlice()
	err = Filter(&dst, &src, func(key interface{}, object interface{}) (skip bool) {
		if v, ok := object.(TFilter); ok {
			if v.ID >= 10 {
				skip = true
			}
		}
		return
	})
	if err != nil {
		t.Fatalf("copy filter() failed: %s", err)
	}
	for _, o := range dst {
		sum += o.ID
	}
	if sum != 45 {
		t.Fatalf("copy filter() failed. Sum is %d", sum)
	}
}

func TestFilterMap(t *testing.T) {
	var (
		err error
		src map[int64]*TFilter
		dst map[int64]*TFilter
		sum int64
	)

	src = createMap()
	err = Filter(&dst, &src, func(key interface{}, object interface{}) (skip bool) {
		if v, ok := object.(TFilter); ok {
			if v.ID >= 10 {
				skip = true
			}
		}
		return
	})
	if err != nil {
		t.Fatalf("copy filter() failed: %s", err)
	}
	for o := range dst {
		sum += o
	}
	if sum != 45 {
		t.Fatal("copy filter() failed")
	}
}
