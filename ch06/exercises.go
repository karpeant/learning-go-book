package main

import (
	"fmt"
)

// exercise 1

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

}

// exercise 2

func UpdateSlice(sl []string, str string) {
	sl[len(sl)-1] = str
	fmt.Println("Slice inside UpdateSlice:", sl)
}

func GrowSlice(sl []string, str string) {
	sl = append(sl, str)
	fmt.Println("Slice inside GrowSlice:", sl)
}

func main() {

	// exercise 1
	newPerson := MakePerson("Bob", "Willimas", 40)
	fmt.Println(newPerson)

	newPersonPointer := MakePersonPointer("Jack", "Willimas", 30)
	fmt.Println(newPersonPointer)

	// exercise 2
	testSlice := []string{"a", "b", "c"}
	fmt.Println("before UpdateSlice", testSlice)
	UpdateSlice(testSlice, "d")
	fmt.Println("after UpdateSlice", testSlice)

	fmt.Println("before GrowSlice", testSlice)
	GrowSlice(testSlice, "e")
	fmt.Println("after GrowSlice", testSlice)

	// exercise 3
	var people []Person

	for i := 0; i < 10000000; i++ {
		people = append(people, MakePerson("Bob", "Willimas", i))
	}

}
