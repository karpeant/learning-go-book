package main

import "fmt"

func exercise_01() {
	var i int = 20
	var f float64 = float64(i)
	fmt.Printf("\nExercise 1:\nthis is int variable i: %d\nthis is float variable f: %f\n", i, f)
}

func exercise_02() {
	const value = 42
	var i int = value
	var f float64 = value
	fmt.Printf("\nExercise 2:\nthis is const variable value assigned to int i: %d\nthis is const variable value assigned to float f: %f\n", i, f)
}

func exercise_03() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI - 1
	bigI = bigI + 1

	fmt.Printf("\nExercise 3:\nthis max byte + 1: %d\nthis is max int32 + 1: %d\nthis is max uint64 + 1: %d\n", b, smallI, bigI)
}

func main() {
	exercise_01()
	exercise_02()
	exercise_03()
}
