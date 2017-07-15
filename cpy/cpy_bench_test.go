package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

type Profile struct {
	Name     string
	Nickname string
	Role     string
	Age      int32
	Years    *int32
	flags    []byte
}

func BenchmarkCopyStruct(b *testing.B) {
	var years int32 = 21
	user := Profile{Name: "Copier lib", Nickname: "Copier lib", Age: 21, Years: &years, Role: "Admin", flags: []byte{'x', 'y', 'z'}}
	for x := 0; x < b.N; x++ {
		_ = All(&Profile{}, &user)
	}
}

func BenchmarkNamaCopy(b *testing.B) {
	var years int32 = 21
	user := Profile{Name: "Copier lib", Nickname: "Copier lib", Age: 21, Years: &years, Role: "Admin", flags: []byte{'x', 'y', 'z'}}
	for x := 0; x < b.N; x++ {
		test := &Profile{
			Name:     user.Name,
			Nickname: user.Nickname,
			Age:      int32(user.Age),
			Years:    user.Years,
		}
		_ = test
	}
}
