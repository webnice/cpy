// Package cpy
package cpy

const tagName = `cpy`

var singleton = &impl{}

// FilterFn Data Filtering Function.
// Return true for skip data
type FilterFn func(key interface{}, object interface{}) (skip bool)
