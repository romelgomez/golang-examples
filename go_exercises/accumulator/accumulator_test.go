package main

import (
	"testing"
)

func Test0(t *testing.T) {
	acc := Accumulator(reducer, 10)

	got15 := acc(5)

	got22 := acc(7)

	got42 := acc(20)

	if got15 != 15 {
		t.Errorf("acc(5) = %d; want 15", got15)
	}

	if got22 != 22 {
		t.Errorf("acc(7) = %d; want 22", got22)
	}

	if got42 != 42 {
		t.Errorf("acc(20) = %d; want 22", got42)
	}

}

func Test1(t *testing.T) {
	acc := Accumulator(reducer, 0)

	got0 := acc(0)

	got02 := acc(0)

	got10 := acc(10)

	if got0 != 0 {
		t.Errorf("acc(0) = %d; want 0", got0)
	}

	if got02 != 0 {
		t.Errorf("acc(0) = %d; want 0", got02)
	}

	if got10 != 10 {
		t.Errorf("acc(10) = %d; want 10", got10)
	}

}

func Test2(t *testing.T) {
	acc := Accumulator(reducer, 0)

	got10 := acc(-10)

	got02 := acc(100)

	got100 := acc(10)

	if got10 != -10 {
		t.Errorf("acc(-10) = %d; want -10", got10)
	}

	if got02 != 90 {
		t.Errorf("acc(100) = %d; want 90", got02)
	}

	if got100 != 100 {
		t.Errorf("acc(10) = %d; want 100", got100)
	}

}

// func reducer(accumulator int, currentValue int) int {
func reducerOdds(first int, second int) int {
	// when given a function that returns the first arg if the second is odd(impar), or the sum of both if the second is even (par)

	if second%2 == 0 {
		return first + second
	} else {
		return first
	}

}

func Test3(t *testing.T) {
	// TODO:

	// ● accumulator › when given a function that returns the first arg if the second is odd, or the sum of both if the second is even › when initial value is 0 › returned function › when called once with odd number › returns the initial value

	// expect(received).resolves.toEqual(expected) // deep equality

	// Expected: 0

	// Received: 3

	acc := Accumulator(reducerOdds, 0)

	result0 := acc(-10)

	result1 := acc(10)

	if result0 != -10 {
		t.Errorf("acc(-10) = %d; want -10", result0)
	}

	if result1 != 0 {
		t.Errorf("acc(10) = %d; want 0", result1)
	}

}

func Test4(t *testing.T) {
	// TODO:

	// ● accumulator › when given a function that returns the first arg if the second is odd, or the sum of both if the second is even › when initial value is 0 › returned function › when called multiple times › returns the number given in each call added (if even) or ignored (if odd) to what the previous call returned

	// expect(received).resolves.toEqual(expected) // deep equality

	// Expected: 938

	// Received: 899
}

func Test5(t *testing.T) {
	// TODO:

	// accumulator › when given a function that returns the first arg if the second is odd, or the sum of both if the second is even › when initial value is 43 › returned function › when called once with even number › returns the number given added to the initial value

	// expect(received).resolves.toEqual(expected) // deep equality

	// Expected: 47

	// Received: 4
}

func Test6(t *testing.T) {
	// TODO:

	// ● accumulator › when given a function that returns the first arg if the second is odd, or the sum of both if the second is even › when initial value is 43 › returned function › when called once with odd number › returns the initial value

	// expect(received).resolves.toEqual(expected) // deep equality

	// Expected: 43

	// Received: 13
}

func Test7(t *testing.T) {
	// TODO:

	// accumulator › when given a function that returns the first arg if the second is odd, or the sum of both if the second is even › when initial value is 43 › returned function › when called multiple times › returns the number given in each call added (if even) or ignored (if odd) to what the previous call returned

	// expect(received).resolves.toEqual(expected) // deep equality

	// Expected: 981

	// Received: 899
}
