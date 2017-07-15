package cpy

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"

// All Сopy everything from one to another
// Копирование полей структур со сравнением по именам полей, не совпадающие поля пропускаются
// Копирование полей с преобразованием pointer -> не pointer (в оба направления)
// Копирование поля в метод и метод в поле при совпадении имён метода и поля (в оба направления),
//  исходный метод должен быть без параметров, результирующи метод должен быть с одним входящим парметром
// Копирование полей с заменой имени поля, описывается в tag структуры
func All(toObj interface{}, fromObj interface{}) error {
	return singleton.Copy(toObj, fromObj)
}

// Only Copy only the listed fields
// Копирование полей структур со сравнением по именам полей, не совпадающие и не перечисленные пропускаются
// Копирование полей с преобразованием pointer -> не pointer (в оба направления)
// Копирование поля в метод и метод в поле при совпадении имён метода и поля (в оба направления),
//  исходный метод должен быть без параметров, результирующи метод должен быть с одним входящим парметром
// Копирование полей с заменой имени поля, описывается в tag структуры
func Selected(toObj interface{}, fromObj interface{}, fields ...string) error {
	//return singleton.Copy(toObj, fromObj)
	return nil
}

// Omit Сopy everything from one to another, but skip listed fields
// Копирование полей структур со сравнением по именам полей, не совпадающие и перечисленные пропускаются
// Копирование полей с преобразованием pointer -> не pointer (в оба направления)
// Копирование поля в метод и метод в поле при совпадении имён метода и поля (в оба направления),
//  исходный метод должен быть без параметров, результирующи метод должен быть с одним входящим парметром
// Копирование полей с заменой имени поля, описывается в tag структуры
func Omit(toObj interface{}, fromObj interface{}, fields ...string) error {
	//return singleton.Copy(toObj, fromObj)
	return nil
}

// FilterByObject Копирование с фильтрацией для всех типов списков (slice, array, map),
// на входе функция возвращающая true или false, вызываемая для каждого объекта данных
// говорящая копировать или пропустить
func FilterByObject(toObj interface{}, fromObj interface{}, filter FilterFn) error {
	//return singleton.Copy(toObj, fromObj)
	return nil
}

// ErrCopyToObjectUnaddressable Error: Copy to object is unaddressable
func ErrCopyToObjectUnaddressable() error { return singleton.ErrCopyToObjectUnaddressable() }

// ErrCopyFromObjectInvalid Error: Copy from object is invalid
func ErrCopyFromObjectInvalid() error { return singleton.ErrCopyFromObjectInvalid() }

// ErrTypeMapNotEqual Error: Type of map is not equal
func ErrTypeMapNotEqual() error { return singleton.ErrTypeMapNotEqual() }
