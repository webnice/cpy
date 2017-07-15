package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import "time"
import "strconv"

type Cat string

func (c Cat) String() string {
	return string(c)
}

type One struct {
	ID                  uint64
	LegoColection       uint32
	HorizontalSeporator uint16
	Place               uint8
	Blocks              uint
	Name                string
	NameInt64           string
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
	Time                string
	Cat                 Cat
}

func (obj *One) String() string {
	return string(obj.Descriptions) + `, name: ` + obj.Name
}

type Two struct {
	NewID    *uint64 `cpy:"name=ID"`
	Name     *string
	Des      []byte
	Complex  string `cpy:"name=String;convert=false"`
	Disabled bool
}

func (obj *Two) Disable(b *bool) {
	if b != nil {
		obj.Disabled = *b
	}
}

type Tm struct {
	Time time.Time
}

func (tm *Tm) Scan(in interface{}) (err error) {
	var value string
	var ok bool
	if value, ok = in.(string); ok {
		tm.Time, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", value)
	}
	return nil
}

type Int64 int64

func (i64 *Int64) Scan(in interface{}) (err error) {
	if value, ok := in.(string); ok {
		var i int64
		i, err = strconv.ParseInt(value, 10, 64)
		*i64 = Int64(i)
	}
	return
}

type Converting struct {
	NewID  int64 `cpy:"name=ID"`
	Des    string
	Int64  Int64 `cpy:"name=NameInt64"`
	Umi    string
	Time   Tm
	String string
	Cat    string
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
		NameInt64:     "-1234567",
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
		Time:          "2017-07-15 02:08:46.691821235 +0000 UTC",
		Cat:           Cat("myau"),
	}
	ret.Marketplace = make([]*string, 2)
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
