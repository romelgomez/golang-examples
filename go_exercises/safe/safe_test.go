package main

import (
	"testing"
)

func Test0(t *testing.T) {
	var moves = []int{10, -5}

	got := FinalValue(moves)

	if got != 5 {
		t.Errorf("FinalValue(moves) = %d; want 5", got)
	}
}

func Test1(t *testing.T) {
	var moves = []int{100, -1, 0}

	got := FinalValue(moves)

	if got != 99 {
		t.Errorf("FinalValue(moves) = %d; want 99", got)
	}
}

func Test3(t *testing.T) {
	var moves = []int{0, 0, 0, 100, 100, 1000, -1000, -10000}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test4(t *testing.T) {
	// returns a number
	var moves = []int{1, 2, 3}

	got := FinalValue(moves)

	if got != 6 {
		t.Errorf("FinalValue(moves) = %d; want 6", got)
	}
}

func Test5(t *testing.T) {
	// when it is not moved ([]) › returns 0
	var moves = []int{}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test6(t *testing.T) {
	// when it receives a single movement › to the right › less than a full turn › returns the same number given
	var moves = []int{93}

	got := FinalValue(moves)

	if got != 93 {
		t.Errorf("FinalValue(moves) = %d; want 93", got)
	}
}

func Test7(t *testing.T) {
	// when it receives a single movement › to the right › one full turn › returns 0
	var moves = []int{100}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test8(t *testing.T) {
	// when it receives a single movement › to the right › more than a full turn › returns what is beyond full turns
	var moves = []int{187}

	got := FinalValue(moves)

	if got != 87 {
		t.Errorf("FinalValue(moves) = %d; want 87", got)
	}
}

func Test10(t *testing.T) {
	var moves = []int{900}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test11(t *testing.T) {
	// when it receives a single movement › to the left › less than a full turn › returns the number plus 100
	var moves = []int{-1}

	got := FinalValue(moves)

	if got != 99 {
		t.Errorf("FinalValue(moves) = %d; want 99", got)
	}
}

func Test12(t *testing.T) {
	// when it receives a single movement › to the left › less than a full turn › returns the number plus 100
	var moves = []int{-83}

	got := FinalValue(moves)

	if got != 17 {
		t.Errorf("FinalValue(moves) = %d; want 17", got)
	}
}

func Test13(t *testing.T) {
	// When it receives a single movement › to the left › one full turn › returns 0
	var moves = []int{-100}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 100", got)
	}
}

func Test14(t *testing.T) {
	// when it receives a single movement › to the left › more than a full turn › returns what is beyond full turns plus 100

	var moves = []int{-183}

	got := FinalValue(moves)

	if got != 17 {
		t.Errorf("FinalValue(moves) = %d; want 17", got)
	}
}

func Test15(t *testing.T) {
	var moves = []int{-800}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test16(t *testing.T) {
	// when it receives multiple movements › all to the right › totalizing less than a full turn › returns final position after all movements
	var moves = []int{1, 43, 25, 2}

	got := FinalValue(moves)

	if got != 71 {
		t.Errorf("FinalValue(moves) = %d; want 71", got)
	}
}

func Test17(t *testing.T) {
	// when it receives multiple movements › all to the right › totalizing one full turn › returns 0

	var moves = []int{21, 30, 49}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test18(t *testing.T) {
	// when it receives multiple movements › all to the right › totalizing multiple full turns › returns 0

	var moves = []int{15, 80, 58, 143, 329, 23, 52}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test19(t *testing.T) {
	// when it receives multiple movements › all to the left › totalizing less than a full turn › returns final position after all movements

	var moves = []int{-4, -43, -22, -6}

	got := FinalValue(moves)

	if got != 25 {
		t.Errorf("FinalValue(moves) = %d; want 25", got)
	}
}

func Test20(t *testing.T) {
	// when it receives multiple movements › all to the left › totalizing one full turn › returns 0

	var moves = []int{-21, -30, -49}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test21(t *testing.T) {
	// when it receives multiple movements › all to the left › totalizing more than a full turn › returns what is beyond full turns plus 100

	var moves = []int{-4, -43, -22, -6, -71, -203}

	got := FinalValue(moves)

	if got != 51 {
		t.Errorf("FinalValue(moves) = %d; want 51", got)
	}
}

func Test22(t *testing.T) {
	// when it receives multiple movements › all to the left › totalizing multiple full turns › returns 0

	var moves = []int{-15, -80, -58, -143, -329, -23, -52}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}
}

func Test23(t *testing.T) {
	// when it receives multiple movements › mixed left and right › totalizing less than a full turn to either side › returns final position after all movements

	var moves = []int{15, -7, 68, 5, -10}

	got := FinalValue(moves)

	if got != 71 {
		t.Errorf("FinalValue(moves) = %d; want 71", got)
	}
}

func Test24(t *testing.T) {
	// when it receives multiple movements › mixed left and right › totalizing one full turn › returns 0

	var moves = []int{74, -53, 79}

	got := FinalValue(moves)

	if got != 0 {
		t.Errorf("FinalValue(moves) = %d; want 0", got)
	}

}

func Test25(t *testing.T) {
	// when it receives multiple movements › mixed left and right › totalizing more than a full turn to either side › returns what is beyond full turns plus 360

	var moves = []int{15, -7, 68, 5, 97, -10}

	got := FinalValue(moves)

	if got != 68 {
		t.Errorf("FinalValue(moves) = %d; want 68", got)
	}
}

func Test(t *testing.T) {
	// when it receives multiple movements › mixed left and right › going through origin from left and right › returns final position after all movements

	var moves = []int{-15, 36, -7, -76, 188}

	got := FinalValue(moves)

	if got != 26 {
		t.Errorf("FinalValue(moves) = %d; want 26", got)
	}
}
