package collections_test

import (
	"fmt"
	"testing"

	"github.com/jeromedoucet/go-common/collections"
)

type stuf struct {
	Id         int
	id         int
	FirstText  string
	SecondText string
	AnInt      int
}

func (s *stuf) String() string {
	return fmt.Sprintf("Id: %d, FirstText: %s, SecondText: %s, AnInt: %d", s.Id, s.FirstText, s.SecondText, s.AnInt)
}

func TestArrayToMapOnIdAttributeWithoutErrorAndWithSlice(t *testing.T) {
	// given
	array := []stuf{
		{Id: 1, FirstText: "Hello"},
		{Id: 2, SecondText: "world"},
		{Id: 3, AnInt: 42},
	}

	// when
	m, err := collections.FromArrayToMap(array, "Id")

	// then
	if err != nil {
		t.Errorf("expect to have no error, but got %s", err.Error())
	}
	if m[1] != array[0] {
		t.Errorf("expect %v to be mapped on index %d, but got %v", array[0], 1, m[1])
	}
	if m[2] != array[1] {
		t.Errorf("expect %v to be mapped on index %d, but got %v", array[1], 2, m[2])
	}
	if m[3] != array[2] {
		t.Errorf("expect %v to be mapped on index %d, but got %v", array[2], 3, m[3])
	}
}

func TestArrayToMapOnIdAttributeWithoutErrorAndWithArray(t *testing.T) {
	// given
	array := [3]stuf{
		{Id: 1, FirstText: "Hello"},
		{Id: 2, SecondText: "world"},
		{Id: 3, AnInt: 42},
	}

	// when
	m, err := collections.FromArrayToMap(array, "Id")

	// then
	if err != nil {
		t.Errorf("expect to have no error, but got %s", err.Error())
	}
	if m[1] != array[0] {
		t.Errorf("expect %v to be mapped on index %d, but got %v", array[0], 1, m[1])
	}
	if m[2] != array[1] {
		t.Errorf("expect %v to be mapped on index %d, but got %v", array[1], 2, m[2])
	}
	if m[3] != array[2] {
		t.Errorf("expect %v to be mapped on index %d, but got %v", array[2], 3, m[3])
	}
}

func TestArrayToMapOnIdAttributeWhenParameterNotAnArrayOrASlice(t *testing.T) {
	// given
	array := "I am not an array"

	// when
	_, err := collections.FromArrayToMap(array, "Id")

	// then
	if err == nil {
		t.Error("expect to have an error, but got nil")
	}
	if _, ok := err.(*collections.TransformationError); !ok {
		t.Errorf("expect the returned error to be an collections.TransformationError, but got %t", err)
	}
	if err.Error() != "Can not handle non array or non slice parameter" {
		t.Errorf("bad error message. got %s", err.Error())
	}
}

func TestArrayToMapOnIdAttributeWhenArrayContainsNoStruct(t *testing.T) {
	// given
	array := []int{1, 2, 3}

	// when
	_, err := collections.FromArrayToMap(array, "Id")

	// then
	if err == nil {
		t.Error("expect to have an error, but got nil")
	}
	if _, ok := err.(*collections.TransformationError); !ok {
		t.Errorf("expect the returned error to be an collections.TransformationError, but got %t", err)
	}
	if err.Error() != "Can only handle array with struct inside" {
		t.Errorf("bad error message. got %s", err.Error())
	}
}

func TestArrayToMapOnIdAttributeWhenUnknownKey(t *testing.T) {
	// given
	array := []stuf{
		{Id: 1, FirstText: "Hello"},
		{Id: 2, SecondText: "world"},
		{Id: 3, AnInt: 42},
	}

	// when
	_, err := collections.FromArrayToMap(array, "UUId")

	// then
	if err == nil {
		t.Error("expect to have an error, but got nil")
	}
	if _, ok := err.(*collections.TransformationError); !ok {
		t.Errorf("expect the returned error to be an collections.TransformationError, but got %t", err)
	}
	if err.Error() != "Unknown key : UUId" {
		t.Errorf("bad error message. got %s", err.Error())
	}
}

func TestArrayToMapOnIdAttributeWhenDuplicatedKey(t *testing.T) {
	// given
	array := []stuf{
		{Id: 1, FirstText: "Hello"},
		{Id: 2, SecondText: "world"},
		{Id: 2, AnInt: 42},
	}

	// when
	_, err := collections.FromArrayToMap(array, "Id")

	// then
	if err == nil {
		t.Error("expect to have an error, but got nil")
	}
	if _, ok := err.(*collections.TransformationError); !ok {
		t.Errorf("expect the returned error to be an collections.TransformationError, but got %t", err)
	}
	if err.Error() != "Duplicated key : 2" {
		t.Errorf("bad error message. got %s", err.Error())
	}
}

func TestArrayToMapOnIdAttributeWhenUnexportedKey(t *testing.T) {
	// given
	array := []stuf{
		{id: 1, FirstText: "Hello"},
		{id: 2, SecondText: "world"},
		{id: 3, AnInt: 42},
	}

	// when
	_, err := collections.FromArrayToMap(array, "id")

	// then
	if err == nil {
		t.Error("expect to have an error, but got nil")
	}
	if _, ok := err.(*collections.TransformationError); !ok {
		t.Errorf("expect the returned error to be an collections.TransformationError, but got %t", err)
	}
	if err.Error() != "Doesn't support unexported key : id" {
		t.Errorf("bad error message. got %s", err.Error())
	}
}

func TestArrayToMapOnIdAttributeWhenEmptyKey(t *testing.T) {
	// given
	array := []stuf{
		{Id: 1, FirstText: "Hello"},
		{Id: 2, SecondText: "world"},
		{Id: 2, AnInt: 42},
	}

	// when
	_, err := collections.FromArrayToMap(array, "")

	// then
	if err == nil {
		t.Error("expect to have an error, but got nil")
	}
	if _, ok := err.(*collections.TransformationError); !ok {
		t.Errorf("expect the returned error to be an collections.TransformationError, but got %t", err)
	}
	if err.Error() != "Empty key" {
		t.Errorf("bad error message. got %s", err.Error())
	}
}
