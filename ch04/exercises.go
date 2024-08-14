package main

import (
	"fmt"
	"math/rand"
)

func exercise01() []int {
	intSlice := make([]int, 0)

	for i := 0; i < 100; i++ {
		randInt := rand.Intn(100)
		intSlice = append(intSlice, randInt)
	}
	return intSlice
}

func exercise02(slice []int) {
	for _, v := range slice {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
			continue
		}
		if v%2 == 0 {
			fmt.Println("Two!")
			continue
		}

		if v%3 == 0 {
			fmt.Println("Three!")
			continue
		}

		fmt.Println("Never mind.")

	}
}

func exercise03() {
	var total int

	for i := 0; i < 10; i++ {
		total = total + 1
		fmt.Println(total)
	}
	fmt.Println(total)
}

func main() {

	sliceExercise01 := exercise01()
	fmt.Println(sliceExercise01)
	exercise02(sliceExercise01)
	exercise03()

}
