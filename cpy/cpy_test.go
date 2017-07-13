package cpy

//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/debug.v1"
import (
	"testing"
)

type One struct {
	ID                  uint64
	LegoColection       uint32
	HorizontalSeporator uint16
	Place               uint8
	Blocks              uint
	Name                string
	Descriptions        []byte `cpy:"name=Des"`
	OnlyPhoto           bool
	Category            int
	Block               int8
	Geo                 int16
	Tables              int32
	Online              int64
	Desktop             float32
	Solutions           float64
	Marketplace         []*string
	ArcMap              map[uint8]string
	Size                [][]int
	Width               []*int8
	Height              []int16
	Umi                 *string
	Disable             *bool
	private             string
}

func (obj *One) String() string {
	return string(obj.Descriptions) + `, name: ` + obj.Name
}

type Two struct {
	Id       *uint64 `cpy:"name=ID"`
	Name     *string
	Des      []byte
	Complex  string `cpy:"name=String"`
	Disabled bool
}

func (obj *Two) Disable(b *bool) {
	if b != nil {
		obj.Disabled = *b
	}
}

func createOne() (ret *One) {
	var nort, west, umi string
	var disable bool

	ret = &One{
		ID:            1,
		LegoColection: 2,
		Place:         3,
		Blocks:        4,
		Name:          "Hello from One.Name",
		Descriptions:  []byte("One.Description"),
		OnlyPhoto:     true,
		Category:      5,
		Block:         -6,
		Geo:           7,
		Tables:        8,
		Online:        9,
		Desktop:       10.003,
		Solutions:     11.1111111,
		Size:          [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		Height:        []int16{128, 64},
		private:       "Private value",
	}
	ret.Marketplace = make([]*string, 2, 2)
	nort, west = "Nort", "West"
	ret.Marketplace[0] = &nort
	ret.Marketplace[1] = &west
	ret.ArcMap = make(map[uint8]string)
	ret.ArcMap[8] = "ArcMap test"
	ret.Width = make([]*int8, 1)
	umi = "Umi"
	ret.Umi = &umi
	disable = true
	ret.Disable = &disable

	return
}

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
