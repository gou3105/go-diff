package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/k0kubun/pp"
)

// DiffResult struct
type DiffResult struct {
	After  map[string]interface{}
	Before map[string]interface{}
}

// Diff struct diff
func Diff(before, after interface{}) error {
	checkStructType(before)

	fmt.Println("")

	checkStructType(after)

	return nil
}

func convertStructToMap(data interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	elem := reflect.TypeOf(data)
	size := elem.NumField()
	fmt.Println(size)
	fmt.Println(reflect.TypeOf(elem))

	return res
}

func convertPtrStructToMap(data interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	fmt.Println(size)
	fmt.Println(reflect.TypeOf(elem))

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		data := elem.Field(i)
		// kind := data.Kind()
		value := data.Interface()
		res[field] = value
	}

	pp.Println(res)

	return res
}

func checkStructType(data interface{}) map[string]interface{} {
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
	G *string
	H Sub
}

type Sub struct {
	A string
}

func main() {
	start := time.Now()
	str := "ptr"
	test1 := TestStruct{
		"test1",
		1,
		[]string{
			"test1",
		},
		1,
		1,
		1.1,
		&str,
		Sub{
			A: "test",
		},
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
		&str,
		Sub{
			A: "test",
		},
	}
	_ = Diff(test1, test2)

	end := time.Now()

	fmt.Printf("%v秒\n", end.Sub(start).Seconds())
}
