package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	ints1 := []float64{8, 4, 10}
	ints2 := []float64{8, 4, 10, 10}
	ints3 := []float64{8, 4, 10, 10.2}

	ints1R := MinDistance(ints1)
	ints2R := MinDistance(ints2)
	ints3R := MinDistance(ints3)

	if ints1R != 2 {
		t.Errorf("MinDistance(ints1) = %d; want 2", ints1R)
	}

	if ints2R != 0 {
		t.Errorf("MinDistance(ints1) = %d; want 0", ints2R)
	}

	if ints3R != 0 {
		t.Errorf("MinDistance(ints1) = %d; want 0", ints3R)
	}
}

func Test0(t *testing.T) {
	// returns a number
	// [1,2,3]

	arr := []float64{1, 2, 3}

	arrR := MinDistance(arr)

	if arrR != 1 {
		t.Errorf("MinDistance(arr) = %d; want 1", arrR)
	}
}

func Test1(t *testing.T) {
	// when it receives a list with two numbers › in descending order › returns the distance between the two numbers as a positive number
	// [10,3]

	arr := []float64{10, 3}

	arrR := MinDistance(arr)

	if arrR != 7 {
		t.Errorf("MinDistance(arr) = %d; want 7", arrR)
	}
}

func Test2(t *testing.T) {
	// when it receives a list with two numbers › in ascending order › returns the distance between the two numbers as a positive number
	// [3,10]

	arr := []float64{3, 10}

	arrR := MinDistance(arr)

	if arrR != 7 {
		t.Errorf("MinDistance(arr) = %d; want 7", arrR)
	}
}

func Test3(t *testing.T) {
	// when it receives a list with many numbers › when at least two numbers in the list are the same › returns 0
	// [1,2,3,4,4,5,6,7]

	arr := []float64{1, 2, 3, 4, 4, 5, 6, 7}

	arrR := MinDistance(arr)

	if arrR != 0 {
		t.Errorf("MinDistance(arr) = %d; want 0", arrR)
	}

}

func Test4(t *testing.T) {
	// when it receives a list with many numbers › when minimum distance is shared between more than one pair of numbers › returns the minimum distance using any of those pairs
	// [9,4,6,5,15,23,22]

	arr := []float64{9, 4, 6, 5, 15, 23, 22}

	arrR := MinDistance(arr)

	if arrR != 1 {
		t.Errorf("MinDistance(arr) = %d; want 1", arrR)
	}

}

func Test5(t *testing.T) {
	// when it receives a list with many numbers › when minimum distance is on two numbers somewhere in the middle of the list › returns the minimum distance correctly
	// [40,10,55,8,32]

	arr := []float64{40, 10, 55, 8, 32}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}
}

func Test6(t *testing.T) {
	// when it receives a list with many numbers › when minimum distance is on the first pair of numbers › returns the minimum distance correctly
	// [10,8,40,55,32]

	arr := []float64{10, 8, 40, 55, 32}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}
}

func Test7(t *testing.T) {
	// when it receives a list with many numbers › when minimum distance is on the last pair of numbers › returns the minimum distance correctly
	// [40,55,32,10,8]

	arr := []float64{40, 55, 32, 10, 8}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}

}

func Test8(t *testing.T) {
	// when it receives a list with many numbers › when minimum distance is a pair of numbers in the middle of the list › returns the minimum distance correctly
	// [40,43,10,8,55,32]

	arr := []float64{40, 43, 10, 8, 55, 32}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}

}

func Test9(t *testing.T) {
	// when it receives a list with many numbers › when minimum distance are from a pair of number at the opposite extremes of the list › returns the minimum distance correctly
	// [10,40,43,55,32,8]

	arr := []float64{10, 40, 43, 55, 32, 8}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}

}

func Test10(t *testing.T) {
	// when it receives a list with many numbers › when the list has negative numbers › returns the minimum distance correctly
	// [-40,10,-55,8,32]

	arr := []float64{-40, 10, -55, 8, 32}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}

}

func Test11(t *testing.T) {
	// when it receives a list with many numbers › when the minimum distance is between two negative numbers › returns the minimum distance correctly
	// [40,-10,55,-8,-32]

	arr := []float64{40, -10, 55, -8, -32}

	arrR := MinDistance(arr)

	if arrR != 2 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}

}

func Test12(t *testing.T) {
	// when it receives a list with many numbers › when the minimum distance is between one negative and one positive number › returns the minimum distance correctly
	// [40,-10,65,8,-32]

	arr := []float64{40, -10, 65, 8, -32}

	arrR := MinDistance(arr)

	if arrR != 18 {
		t.Errorf("MinDistance(arr) = %d; want 2", arrR)
	}
}
