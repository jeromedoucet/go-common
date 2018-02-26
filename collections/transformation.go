package collections

import (
	"fmt"
	"reflect"
	"unicode"
)

// here is some functions that perform
// transformations over collections (array, slice and map)

// TransformationError is raised when
// an issue happened during a transformation
type TransformationError struct {
	msg string
}

func (e TransformationError) Error() string {
	return e.msg
}

func newTransformationError(msg string) *TransformationError {
	e := new(TransformationError)
	e.msg = msg
	return e
}

// FromArrayToMap transform an arbitrary array or slice to a map.
// using a field of array elements. If the array doesn't contains
// struct, a nil map and a non nil error are returned.
//
// the first argument must be the slice or the array. If its
// type doesn't match, a nil map and a non nil error are returned.
//
// the second argument is the name of the field to used as key for the returned map
// This function only support non empty and exported (public) field. If conditions
// are not satisfied, a nil map and a non nil error are returned.
//
// This function check also that there is no key duplication. Such behavior
// avoid silently loosing data, leading to have buggy programs. Like for the
// others rules, if a duplication is detected, a nil map and a non nil error are returned.
//
// NOTE : This function has been designed to be used in tests functions, not for production code.
// It is convenient to avoid pollute tests code with loop when a map is needed from an array, but
// it is not designed neither for performance nor for being precise on typing.
func FromArrayToMap(array interface{}, fieldName string) (map[interface{}]interface{}, error) {
	if len(fieldName) == 0 {
		return nil, newTransformationError("Empty key")
	}
	if unicode.IsLower([]rune(fieldName)[0]) {
		return nil, newTransformationError(fmt.Sprintf("Doesn't support unexported key : %s", fieldName))
	}
	val := reflect.ValueOf(array)
	t := val.Type()
	if t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return nil, newTransformationError("Can not handle non array or non slice parameter")
	}
	res := make(map[interface{}]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		if elem.Kind() != reflect.Struct {
			return nil, newTransformationError("Can only handle array with struct inside")
		}
		f := elem.FieldByName(fieldName)
		if !f.IsValid() {
			return nil, newTransformationError(fmt.Sprintf("Unknown key : %s", fieldName))
		}
		_, alreadyExist := res[f.Interface()]
		if alreadyExist {
			return nil, newTransformationError(fmt.Sprintf("Duplicated key : %v", f.Interface()))
		}
		res[f.Interface()] = elem.Interface()
	}
	return res, nil
}
