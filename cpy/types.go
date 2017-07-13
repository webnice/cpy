package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import ()

const tagName = `cpy`

var singleton = &impl{}

// impl is an implementation of package
type impl struct {
}

// FilterFn Функция фильтрации данных
type FilterFn func(object interface{}) bool
