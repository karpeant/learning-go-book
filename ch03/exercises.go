package main

import "fmt"

func exercise01() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subslice1 := make([]string, 2)
	subslice2 := make([]string, 3)
	subslice3 := make([]string, 2)

	copy(subslice1, greetings[:2])
	copy(subslice2, greetings[1:4])
	copy(subslice3, greetings[3:])

	fmt.Println(subslice1, subslice2, subslice3)
}

func exercise02() {
	message := "Hi 👩 and 👨"

	runeMessage := []rune(message)

	rune4thAsStr := string(runeMessage[3])

	fmt.Println(rune4thAsStr)
}

func exercise03() {

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	struck1 := Employee{
		"firstName1",
		"lastName1",
		1,
	}
	struck2 := Employee{
		firstName: "firstName2",
		lastName:  "lastName2",
		id:        2,
	}
	struck3 := Employee{}
	struck3.firstName = "firstName3"
	struck3.lastName = "lastName3"
	struck3.id = 3

	fmt.Println(struck1, struck2, struck3)
}

func main() {
	exercise01()
	exercise02()
	exercise03()
}
