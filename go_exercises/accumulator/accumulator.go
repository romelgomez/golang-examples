package main

func Accumulator(reducer func(int, int) int, init int) func(int) int {
	var acc int = init

	return func(i int) int {
		acc = reducer(acc, i)

		return acc
	}
}

func reducer(accumulator int, currentValue int) int {
	return accumulator + currentValue
}

func main() {

}
