package main

import (
	"fmt"

	"github.com/amesinp/learning_go/traversy_crash/strutil"
)

func greet(name string) string {
	return "Hello, " + name
}

func printReverse(str string) {
	fmt.Println(strutil.Reverse(str))
}

func main() {
	var name = "Precious"
	var age int32 = 21
	const userType = "None"

	// Shorthand syntax
	email, isNew := "amesinlolaprecious@gmail.com", true

	fmt.Println(name, age, userType, email, isNew)
	fmt.Println(greet(name))

	printReverse("hannah")

	// Arrays

	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	colourArr := [2]string{"Red", "White"}

	colourSlice := []string{"Red", "White", "Black", "Blue"}

	fmt.Println(fruitArr)
	fmt.Println(colourArr)
	fmt.Println(colourSlice)

	if age < 18 {
		fmt.Println(name + " is a minor")
	} else {
		fmt.Println(name + " is an adult")
	}

	switch userType {
	case "Admin":
		fmt.Println("Welcome to the admin panel")
	case "User":
		fmt.Println("Welcome to the user panel")
	default:
		fmt.Println("Welcome to the default panel")
	}

	// Loop

	for i := 1; i <= 3; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// Long method
	x := 1
	for x <= 3 {
		fmt.Println(x)
		x++
	}

	// Maps

	numbers := make(map[int]string)

	numbers[1] = "One"
	numbers[2] = "Two"

	fmt.Println(numbers)

	anotherNumbers := map[int]string{1: "One", 2: "Two", 3: "Three"}

	fmt.Println(anotherNumbers)
	delete(anotherNumbers, 1)
	fmt.Println(anotherNumbers)

	// Range

	for i, colour := range colourSlice {
		fmt.Printf("I = %d, Colour = %s\n", i, colour)
	}

	for _, colour := range colourSlice {
		fmt.Printf("Colour = %s\n", colour)
	}

	for key, value := range anotherNumbers {
		fmt.Printf("Digit = %d, Word = %s\n", key, value)
	}

	for _, value := range anotherNumbers {
		fmt.Printf("Word = %s\n", value)
	}

	// Pointers

	a := 100
	b := &a

	fmt.Printf("typeof A: %T, typeof B: %T\n", a, b)

	*b = 200

	fmt.Println(a)

}
