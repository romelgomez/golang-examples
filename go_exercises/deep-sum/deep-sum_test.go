package main

import (
	"encoding/json"
	"testing"
)

func ToString(i []interface{}) string {

	mJson, err := json.Marshal(i)

	if err != nil {
		return "error"
	}

	return string(mJson)
}

func Test0(t *testing.T) {

	//
	//
	// [[[[[10]]],{"a":10,"b":[10,10]},10,[{"c":10},10,[10,10,10,[10,{"a":10,"b":10,"c":10,"d":{"a":10,"b":[10]}},10]]]]]
	//
	//

	iten0 := []interface{}{[]interface{}{[]interface{}{10}}}

	iten1 := map[string]interface{}{
		"a": 10,
		"c": []interface{}{10, 10},
	}

	item3 := map[string]interface{}{
		"c": 10,
	}

	item4 := map[string]interface{}{
		"a": 10,
		"b": 10,
		"c": 10,
		"d": map[string]interface{}{
			"a": 10,
			"b": []interface{}{10},
		},
	}

	arr := []interface{}{[]interface{}{iten0, iten1, 10, []interface{}{item3, 10, []interface{}{10, 10, 10, []interface{}{10, item4, 10}}}}}

	got := DeepSum(arr)

	if got != 170 {
		t.Errorf("DeepSum(arr) = %d; want 170", got)
	}

}

func Test1(t *testing.T) {

	// slice
	// map

	var mapA map[string]interface{}
	var mapB map[string]interface{}

	mapA = map[string]interface{}{
		"a": 10,
		"b": 10,
		"c": []interface{}{10, 10},
	}

	mapB = map[string]interface{}{
		"a": 10,
		"b": 10,
		"c": []interface{}{10, 10, mapA},
	}

	arr0 := []interface{}{10, 10, []interface{}{10, 10, []interface{}{10, 10}}, mapB}
	arr1 := []interface{}{10, 10, []interface{}{10, 10}, mapA}
	arr2 := []interface{}{10, 10, []interface{}{10, 10}, 10}
	arr3 := []interface{}{10, 10, mapA}
	arr4 := []interface{}{10, 10}

	acc0 := DeepSum(arr0)
	acc1 := DeepSum(arr1)
	acc2 := DeepSum(arr2)
	acc3 := DeepSum(arr3)
	acc4 := DeepSum(arr4)

	if acc0 != 140 {
		t.Errorf("DeepSum(arr) = %d; want 140", acc0)
	}

	if acc1 != 80 {
		t.Errorf("DeepSum(arr) = %d; want 80", acc1)
	}

	if acc2 != 50 {
		t.Errorf("DeepSum(arr) = %d; want 50", acc2)
	}

	if acc3 != 60 {
		t.Errorf("DeepSum(arr) = %d; want 60", acc3)
	}

	if acc4 != 20 {
		t.Errorf("DeepSum(arr) = %d; want 20", acc4)
	}

}

func Test2(t *testing.T) {
	// when the argument is an empty list › returns a number

	arr := []interface{}{}

	acc0 := DeepSum(arr)

	if acc0 != 0 {
		t.Errorf("DeepSum(arr) = %d; want 0", acc0)
	}

}

func Test4(t *testing.T) {
	// when the argument is an empty list › returns 0

	arr := []interface{}{}

	acc0 := DeepSum(arr)

	if acc0 != 0 {
		t.Errorf("DeepSum(arr) = %d; want 0", acc0)
	}
}

func Test5(t *testing.T) {
	// when the argument is just numbers › returns a number
	// when the argument is just numbers › returns the sum of all the numbers

	arr := []interface{}{4, 13, 1, 7}

	acc0 := DeepSum(arr)

	if acc0 != 25 {
		t.Errorf("DeepSum(arr) = %d; want 25", acc0)
	}
}

func Test7(t *testing.T) {
	// when the argument is just one level of array of numbers › returns a number
	// when the argument is just one level of array of numbers › returns the sum of all the numbers

	arr := []interface{}{[]interface{}{4}, []interface{}{13, 1}, []interface{}{7}}

	acc0 := DeepSum(arr)

	if acc0 != 25 {
		t.Errorf("DeepSum(arr) = %d; want 25", acc0)
	}

	arrStr := ToString(arr)

	if arrStr != "[[4],[13,1],[7]]" {
		t.Errorf("ToString(arr) = %s; want [[4],[13,1],[7]]", arrStr)
	}
}

func Test10(t *testing.T) {
	// when the argument is just one level of objects with number values › returns a number
	// when the argument is just one level of objects with number values › returns the sum of all the numbers
	// when the argument is just one level of objects with number values › returns the sum of all the numbers
	// [{"a":4},{"b":13,"c":1},{"a":7}]

	mapA := map[string]interface{}{
		"a": 4,
	}

	mapB := map[string]interface{}{
		"b": 13,
		"c": 1,
	}

	mapC := map[string]interface{}{
		"a": 7,
	}

	arr := []interface{}{mapA, mapB, mapC}

	acc0 := DeepSum(arr)

	if acc0 != 25 {
		t.Errorf("DeepSum(arr) = %d; want 25", acc0)
	}

	arrStr := ToString(arr)

	if arrStr != "[{\"a\":4},{\"b\":13,\"c\":1},{\"a\":7}]" {
		t.Errorf("ToString(arr) = %s; want [[4],[13,1],[7]]", arrStr)
	}
}

func Test12(t *testing.T) {
	// when the argument is one level of mixed numbers, arrays of numbers and objects with number values › returns a number
	// when the argument is one level of mixed numbers, arrays of numbers and objects with number values › returns the sum of all the numbers
	// [{\"a\":4},[13,1],7]

	mapA := map[string]interface{}{
		"a": 4,
	}

	arr := []interface{}{mapA, []interface{}{13, 1}, 7}

	acc0 := DeepSum(arr)

	if acc0 != 25 {
		t.Errorf("DeepSum(arr) = %d; want 25", acc0)
	}

	arrStr := ToString(arr)

	if arrStr != "[{\"a\":4},[13,1],7]" {
		t.Errorf("ToString(arr) = %s; want [{\"a\":4},[13,1],7]", arrStr)
	}

}

func Test15(t *testing.T) {

	// when the argument is nested arrays, objects and plain numbers in multiple levels and combinations › returns the sum of all the numbers
	// [[[[4]]],{\"a\":13,\"b\":[6,4]},1,[{\"c\":7},5,[3,9,8,[5,{\"a\":2,\"b\":10,\"c\":3,\"d\":{\"a\":4,\"b\":[6]}},7]]]]

	iten0 := []interface{}{[]interface{}{[]interface{}{4}}}

	iten1 := map[string]interface{}{
		"a": 13,
		"b": []interface{}{6, 4},
	}

	item3 := map[string]interface{}{
		"c": 7,
	}

	item4 := map[string]interface{}{
		"a": 2,
		"b": 10,
		"c": 3,
		"d": map[string]interface{}{
			"a": 4,
			"b": []interface{}{6},
		},
	}

	arr := []interface{}{iten0, iten1, 1, []interface{}{item3, 5, []interface{}{3, 9, 8, []interface{}{5, item4, 7}}}}

	got := DeepSum(arr)

	if got != 97 {
		t.Errorf("DeepSum(arr) = %d; want 97", got)
	}

	arrStr := ToString(arr)

	if arrStr != "[[[[4]]],{\"a\":13,\"b\":[6,4]},1,[{\"c\":7},5,[3,9,8,[5,{\"a\":2,\"b\":10,\"c\":3,\"d\":{\"a\":4,\"b\":[6]}},7]]]]" {
		t.Errorf("ToString(arr) = %s; want [[[[4]]],{\"a\":13,\"b\":[6,4]},1,[{\"c\":7},5,[3,9,8,[5,{\"a\":2,\"b\":10,\"c\":3,\"d\":{\"a\":4,\"b\":[6]}},7]]]]", arrStr)
	}

}
