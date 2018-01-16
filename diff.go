package main

import (
	"fmt"
	"reflect"
)

// DiffResult struct
type DiffResult struct {
	After  map[string]interface{}
	Before map[string]interface{}
}

// Diff struct diff
func Diff(before, after interface{}) error {
	checkType(before)

	fmt.Println("")

	checkType(after)

	return nil
}

func convertStructToMap(data interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	elem := reflect.TypeOf(data)
	size := elem.NumField()
	fmt.Println(size)

	return res
}

func convertPtrStructToMap(data interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	fmt.Println(size)

	return res
}

func checkType(data interface{}) map[string]interface{} {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Struct:
		return convertStructToMap(data)
	case reflect.Ptr:
		return convertPtrStructToMap(data)
	}
	return nil
}

type TestStruct struct {
	A string
	B int
	C []string
	D int16
	E int32
	F float32
}

func main() {
	test1 := TestStruct{
		"test1",
		1,
		[]string{
			"test1",
		},
		1,
		1,
		1.1,
	}
	test2 := &TestStruct{
		"test2",
		2,
		[]string{
			"test2",
		},
		2,
		2,
		2.2,
	}
	_ = Diff(test1, test2)
}
