// Package cpy
package cpy

import "testing"

type Profile struct {
	Name     string
	Nickname string
	Role     string
	Age      int32
	Years    *int32
	flags    []byte
}

func BenchmarkCopyStruct(b *testing.B) {
	var (
		years int32
		user  Profile
		n     int
	)

	years = 21
	user = Profile{Name: "Copier lib", Nickname: "Copier lib", Age: 21, Years: &years, Role: "Admin", flags: []byte{'x', 'y', 'z'}}
	for n = 0; n < b.N; n++ {
		_ = All(&Profile{}, &user)
	}
}

func BenchmarkNamaCopy(b *testing.B) {
	var (
		years int32
		user  Profile
		n     int
	)

	years = 21
	user = Profile{Name: "Copier lib", Nickname: "Copier lib", Age: 21, Years: &years, Role: "Admin", flags: []byte{'x', 'y', 'z'}}
	for n = 0; n < b.N; n++ {
		test := &Profile{
			Name:     user.Name,
			Nickname: user.Nickname,
			Age:      user.Age,
			Years:    user.Years,
		}
		_ = test
	}
}
