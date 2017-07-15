package cpy_test

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/webnice/cpy.v1/cpy"
)

// Copying all data without filtration with type conversion and field matching.
func ExampleAll_everything() {
	// Src Source structure
	type Src struct {
		ID       int64
		Name     string
		Value    []byte `cpy:"name=Description"` // Overriding the field name to match the field in the destination structure
		CreateAt func() string
	}

	// Dst Destination structure
	type Dst struct {
		MyID        int     `cpy:"name=ID"   json:"id"`    // Overriding the field name to match the field name in the data source structure
		Title       *string `cpy:"name=Name" json:"title"` // Overriding the field name to match the field name in the data source structure
		Description string  `                json:"des"`
		CreateAt    string  `                json:"crateAt"`
	}

	var source []Src
	var destination []*Dst

	fn := func() string {
		return time.Date(2017, 7, 15, 10, 35, 24, 0, time.UTC).String()
	}

	source = []Src{
		Src{ID: 1, Name: "Aiden", Value: []byte("Smith"), CreateAt: fn},
		Src{ID: 2, Name: "Liam", Value: []byte("Johnson"), CreateAt: fn},
		Src{ID: 3, Name: "Isabella", Value: []byte("Brown"), CreateAt: fn},
	}
	// Сopy everything from one to another
	err := cpy.All(&destination, &source)
	if err != nil {
		log.Fatalf("Error copy: %s", err.Error())
	}
	// Output result
	b, _ := json.MarshalIndent(destination, "", "  ")
	fmt.Printf("%s\n", string(b))

	// Output:
	// [
	//   {
	//     "id": 1,
	//     "title": "Aiden",
	//     "des": "Smith",
	//     "crateAt": "2017-07-15 10:35:24 +0000 UTC"
	//   },
	//   {
	//     "id": 2,
	//     "title": "Liam",
	//     "des": "Johnson",
	//     "crateAt": "2017-07-15 10:35:24 +0000 UTC"
	//   },
	//   {
	//     "id": 3,
	//     "title": "Isabella",
	//     "des": "Brown",
	//     "crateAt": "2017-07-15 10:35:24 +0000 UTC"
	//   }
	// ]
}
