package cpy_test

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/webnice/cpy.v1/cpy"
)

func ExampleAll() {
	type Src struct {
		ID    int64
		Name  string
		Value []byte `cpy:"name=Description"` // Overriding the field name to match the field in the destination structure
	}

	type Dst struct {
		MyID        int     `cpy:"name=ID"`   // Overriding the field name to match the field name in the data source structure
		Title       *string `cpy:"name=Name"` // Overriding the field name to match the field name in the data source structure
		Description string
	}

	var source []Src
	var destination []*Dst

	source = []Src{
		Src{ID: 1, Name: "Aiden", Value: []byte("Smith")},
		Src{ID: 2, Name: "Liam", Value: []byte("Johnson")},
		Src{ID: 3, Name: "Isabella", Value: []byte("Brown")},
	}
	if err := cpy.All(&destination, &source); err != nil {
		log.Fatalf("Error copy: %s", err.Error())
	}
	b, _ := json.MarshalIndent(destination, "", "  ")
	fmt.Printf("%s\n", string(b))

	// Output:
	// [
	//   {
	//     "MyID": 1,
	//     "Title": "Aiden",
	//     "Description": "Smith"
	//   },
	//   {
	//     "MyID": 2,
	//     "Title": "Liam",
	//     "Description": "Johnson"
	//   },
	//   {
	//     "MyID": 3,
	//     "Title": "Isabella",
	//     "Description": "Brown"
	//   }
	// ]
}
