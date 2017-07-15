package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"

const tagName = `cpy`

var singleton = &impl{}

// FilterFn Data Filtering Function.
// Return true for skip data
type FilterFn func(key interface{}, object interface{}) (skip bool)
