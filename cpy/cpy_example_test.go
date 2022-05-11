// Package cpy_test
package cpy_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/webnice/kit/module/cfg/cpy"
)

// Copying all data without filtration with type conversion and field matching
func ExampleAll_everything() {
	type (
		// Src Source structure
		Src struct {
			ID       int64
			Name     string
			Value    []byte `cpy:"name=Description"` // Overriding the field name to match the field in the destination structure
			CreateAt func() string
		}
		// Dst Destination structure
		Dst struct {
			MyID        int     `cpy:"name=ID"   json:"id"`    // Overriding the field name to match the field name in the data source structure
			Title       *string `cpy:"name=Name" json:"title"` // Overriding the field name to match the field name in the data source structure
			Description string  `                json:"des"`
			CreateAt    string  `                json:"crateAt"`
		}
	)
	var (
		err         error
		fn          func() string
		source      []Src
		destination []*Dst
		b           []byte
	)

	fn = func() string { return time.Date(2017, 7, 15, 10, 35, 24, 0, time.UTC).String() }
	source = []Src{
		{ID: 1, Name: "Aiden", Value: []byte("Smith"), CreateAt: fn},
		{ID: 2, Name: "Liam", Value: []byte("Johnson"), CreateAt: fn},
		{ID: 3, Name: "Isabella", Value: []byte("Brown"), CreateAt: fn},
	}
	// Copy everything from one to another
	if err = cpy.All(&destination, &source); err != nil {
		log.Fatalf("copy, error: %s", err)
	}
	// Output result
	b, _ = json.MarshalIndent(destination, "", "  ")
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

// Copying slice with filtration function
func ExampleFilter_filtration() {
	type (
		// Src Source structure
		Src struct {
			ID       int64
			FullName string
			Age      int32
		}
		// Dst Destination structure
		Dst struct {
			NewID int     `cpy:"name=ID"       json:"id"`    // Overriding the field name to match the field name in the data source structure
			Title *string `cpy:"name=FullName" json:"title"` // Overriding the field name to match the field name in the data source structure
		}
	)
	var (
		err         error
		source      []*Src
		destination []Dst
		b           []byte
	)

	source = []*Src{
		{ID: 1, FullName: "Aiden Smith", Age: 17},
		{ID: 2, FullName: "Liam Johnson", Age: 19},
		{ID: 3, FullName: "Isabella Brown", Age: 21},
	}
	// Copy everything from one to another
	err = cpy.Filter(&destination, &source, func(key interface{}, object interface{}) (skip bool) {
		skip = true // By default, all rows are skipped
		// This is filtration function
		// key is index of value in slice and key in map (In this example is not required)
		//
		// Filtering function always receives a copy of the object, regardless of how slice is defined
		// Therefore, always lead to the type of the slice element (Src), not a (*Src)
		if v, ok := object.(Src); ok {
			// filter by age >= 18
			if v.Age >= 18 {
				skip = false
			}
		}

		return
	})
	if err != nil {
		log.Fatalf("copy, error: %s", err)
	}
	// Output result
	b, _ = json.MarshalIndent(destination, "", "  ")
	fmt.Printf("%s\n", string(b))

	// Output:
	// [
	//   {
	//     "id": 2,
	//     "title": "Liam Johnson"
	//   },
	//   {
	//     "id": 3,
	//     "title": "Isabella Brown"
	//   }
	// ]
}

// Copying selected fields of structures
func ExampleSelect_byField() {
	type (
		// MyType Source and destination structure
		MyType struct {
			ID          int64   `json:"id"`
			FullName    string  `json:"name"`
			Age         int32   `json:"age"`
			Description string  `json:"des"`
			Comments    *string `json:"-"`
		}
	)
	var (
		err         error
		source      []*MyType
		destination []MyType
		b           []byte
	)

	source = []*MyType{
		{ID: 1, FullName: "Aiden Smith", Age: 17, Description: "User Aiden Smith"},
		{ID: 2, FullName: "Liam Johnson", Age: 19, Description: "User Liam Johnson"},
		{ID: 3, FullName: "Isabella Brown", Age: 21, Description: "User Isabella Brown"},
	}
	// Copy only ID and FullName fields
	if err = cpy.Select(&destination, &source, "ID", "FullName"); err != nil {
		log.Fatalf("copy, error: %s", err)
	}
	// Output result
	b, _ = json.MarshalIndent(destination, "", "  ")
	fmt.Printf("%s\n", string(b))

	// Output:
	// [
	//   {
	//     "id": 1,
	//     "name": "Aiden Smith",
	//     "age": 0,
	//     "des": ""
	//   },
	//   {
	//     "id": 2,
	//     "name": "Liam Johnson",
	//     "age": 0,
	//     "des": ""
	//   },
	//   {
	//     "id": 3,
	//     "name": "Isabella Brown",
	//     "age": 0,
	//     "des": ""
	//   }
	// ]
}

// Copying all fields of structures, but skip listed fields
func ExampleOmit_byField() {
	type (
		// MyType Source and destination structure
		MyType struct {
			ID          int64  `json:"id"`
			FullName    string `json:"name"`
			Age         int32  `json:"age"`
			Description string `json:"des"`
		}
	)
	var (
		err         error
		source      []MyType
		destination []*MyType
		b           []byte
	)

	source = []MyType{
		{ID: 1, FullName: "Aiden Smith", Age: 17, Description: "User Aiden Smith"},
		{ID: 2, FullName: "Liam Johnson", Age: 19, Description: "User Liam Johnson"},
		{ID: 3, FullName: "Isabella Brown", Age: 21, Description: "User Isabella Brown"},
	}
	// Skip Description field
	if err = cpy.Omit(&destination, &source, "Description"); err != nil {
		log.Fatalf("copy, error: %s", err)
	}
	// Output result
	b, _ = json.MarshalIndent(destination, "", "  ")
	fmt.Printf("%s\n", string(b))

	// Output:
	// [
	//   {
	//     "id": 1,
	//     "name": "Aiden Smith",
	//     "age": 17,
	//     "des": ""
	//   },
	//   {
	//     "id": 2,
	//     "name": "Liam Johnson",
	//     "age": 19,
	//     "des": ""
	//   },
	//   {
	//     "id": 3,
	//     "name": "Isabella Brown",
	//     "age": 21,
	//     "des": ""
	//   }
	// ]
}
