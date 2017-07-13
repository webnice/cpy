package cpy

//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/debug.v1"
import (
	"database/sql"
	"reflect"
	"strings"
)

func init() {
	debug.Nop()
}

// Сopy everything
func (cpy *impl) Copy(toObj interface{}, fromObj interface{}) (err error) {
	var from, to, src, dst reflect.Value
	var fromT, toT reflect.Type
	var isSlice bool
	var i, size int

	// Values
	to, from = cpy.Indirect(reflect.ValueOf(toObj)), cpy.Indirect(reflect.ValueOf(fromObj))
	if isSlice, size, err = cpy.Check(to, from); err != nil {
		return
	}

	// If possible to assign
	if from.Type().AssignableTo(to.Type()) {
		to.Set(from)
		return
	}
	// Types
	fromT, toT = cpy.IndirectType(from.Type()), cpy.IndirectType(to.Type())

	// The magic :)
	for i = 0; i < size; i++ {
		if isSlice {
			if from.Kind() == reflect.Slice {
				src = cpy.Indirect(from.Index(i))
			} else {
				src = cpy.Indirect(from)
			}
			dst = cpy.Indirect(reflect.New(toT).Elem())
		} else {
			src = cpy.Indirect(from)
			dst = cpy.Indirect(to)
		}
		// Copy from method to field
		if err = cpy.CopyFromMethod(toT, fromT, dst, src); err != nil {
			return
		}
		// Copy from field to field or method
		if err = cpy.CopyFromField(toT, fromT, dst, src); err != nil {
			return
		}
		if isSlice {
			switch {
			case dst.Addr().Type().AssignableTo(to.Type().Elem()):
				to.Set(reflect.Append(to, dst.Addr()))
			case dst.Type().AssignableTo(to.Type().Elem()):
				to.Set(reflect.Append(to, dst))
			}
		}
	}

	return
}

// CopyFromField Copy from field to field or method
func (cpy *impl) CopyFromField(toT reflect.Type, fromT reflect.Type, dst reflect.Value, src reflect.Value) (err error) {
	const paramName = `name`
	var fromF reflect.Value
	var field reflect.StructField
	var srcName, dstName string

	// Copy from field or method to field
	for _, field = range cpy.Fields(fromT) {
		srcName = field.Name
		if dstName = cpy.FieldReplaceName(field, paramName); dstName == "" {
			dstName = srcName
		}
		fromF = src.FieldByName(srcName)
		if fromF.IsValid() {
			if err = cpy.SetToFieldOrMethod(dst, dstName, fromF, srcName); err != nil {
				return
			}
		}
	}

	return
}

// Copy from method to field
func (cpy *impl) CopyFromMethod(toT reflect.Type, fromT reflect.Type, dst reflect.Value, src reflect.Value) (err error) {
	const paramName = `name`
	var fromM reflect.Value
	var field reflect.StructField
	var srcName, dstName string

	for _, field = range cpy.Fields(toT) {
		dstName = field.Name
		if srcName = cpy.FieldReplaceName(field, paramName); srcName == "" {
			srcName = dstName
		}
		if src.CanAddr() {
			fromM = src.Addr().MethodByName(srcName)
		} else {
			fromM = src.MethodByName(srcName)
		}
		if fromM.IsValid() {
			if err = cpy.SetToFieldOrMethod(dst, dstName, fromM, srcName); err != nil {
				return
			}
		}
	}

	return
}

// SetToFieldOrMethod Set value to field or method
func (cpy *impl) SetToFieldOrMethod(dst reflect.Value, dstName string, from reflect.Value, srcName string) (err error) {
	const paramName = `name`
	var toF, toM, fromM reflect.Value
	var field reflect.StructField
	var values []reflect.Value
	var name string

	// Запрос по имени поля
	toF = dst.FieldByName(dstName)
	// Поиск по тегу
	if !toF.IsValid() {
		for _, field = range cpy.Fields(dst.Type()) {
			if name = field.Name; cpy.FieldReplaceName(field, paramName) == dstName {
				toF = dst.FieldByName(name)
				break
			}
		}
	}

	// Если field
	if toF.IsValid() {
		// Try to can set
		if toF.CanSet() {
			if !cpy.Set(toF, from) {
				if from.Kind() == reflect.Func &&
					from.Type().NumIn() == 0 &&
					from.Type().NumOut() >= 1 {
					if values = from.Call([]reflect.Value{}); len(values) > 0 {
						cpy.Set(toF, values[0])
					}
				} else {
					if err = cpy.Copy(toF.Addr().Interface(), from.Interface()); err != nil {
						return
					}
				}
			}
		} else {
			if from.CanAddr() {
				fromM = from.Addr().MethodByName(srcName)
			} else {
				fromM = from.MethodByName(srcName)
			}
			values = fromM.Call([]reflect.Value{})
			if len(values) > 0 {
				cpy.Set(toF, values[0])
			}
		}
	} else {
		// Try to set call method
		if dst.CanAddr() {
			toM = dst.Addr().MethodByName(dstName)
		} else {
			toM = dst.MethodByName(dstName)
		}
		if toM.IsValid() &&
			toM.Type().NumIn() == 1 &&
			from.Type().AssignableTo(toM.Type().In(0)) {
			toM.Call([]reflect.Value{from})
		}
	}

	return
}

// Indirect value get
func (cpy *impl) Indirect(rv reflect.Value) reflect.Value {
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return rv
}

// Indirect type get
func (cpy *impl) IndirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

// Checks if input objects is correct
func (cpy *impl) Check(to reflect.Value, from reflect.Value) (isSlice bool, size int, err error) {
	if !from.IsValid() {
		err = cpy.ErrCopyFromObjectInvalid()
	}
	if !to.CanAddr() {
		err = cpy.ErrCopyToObjectUnaddressable()
	}
	if err != nil {
		return
	}
	size = 1
	if to.Kind() == reflect.Slice {
		isSlice = true
		if from.Kind() == reflect.Slice {
			size = from.Len()
		}
	}
	return
}

// Set value
func (cpy *impl) Set(to reflect.Value, from reflect.Value) (ok bool) {
	var scanner sql.Scanner

	if from.IsValid() {
		if to.Kind() == reflect.Ptr {
			if to.IsNil() {
				to.Set(reflect.New(to.Type().Elem()))
			}
			to = to.Elem()
		}
		if from.Type().ConvertibleTo(to.Type()) {
			to.Set(from.Convert(to.Type()))
		} else if scanner, ok = to.Addr().Interface().(sql.Scanner); ok {
			_ = scanner.Scan(from.Interface())
		} else if from.Kind() == reflect.Ptr {
			ok = cpy.Set(to, from.Elem())
		}
	}

	return
}

// Fields to StructField
func (cpy *impl) Fields(rt reflect.Type) (ret []reflect.StructField) {
	var i int
	var v reflect.StructField

	if rt = cpy.IndirectType(rt); rt.Kind() == reflect.Struct {
		for i = 0; i < rt.NumField(); i++ {
			v = rt.Field(i)
			if v.Anonymous {
				ret = append(ret, cpy.Fields(v.Type)...)
				continue
			}
			ret = append(ret, v)
		}
	}

	return
}

// FieldReplaceName Get field name from tag
func (cpy *impl) FieldReplaceName(field reflect.StructField, name string) (ret string) {
	var tag string
	var params, tmp []string
	var i int

	if tag = field.Tag.Get(tagName); tag == "" {
		return
	}
	params = strings.Split(tag, ";")
	for i = range params {
		if tmp = strings.Split(params[i], "="); len(tmp) > 1 {
			if tmp[0] == name {
				ret = tmp[1]
			}
		}
	}

	return
}
