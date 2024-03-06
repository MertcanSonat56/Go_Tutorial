package main

import "fmt"

func main() {

	// Variables and Constants
	// Constants are like variables, except that their values cannot be changed
	var nickname string = "WaweX56"
	const conferenceTickets int = 50

	fmt.Println(nickname)
	fmt.Println(conferenceTickets)

	// Data Types
	var firstName string
	var lastName string
	var email string

	// Pointer: pointer is a variable that points to the memory address of another variable
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Printf("firstName: %v , lastName: %v , email: %v", firstName, lastName, email)

	// Arrays & Slices
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	fmt.Println("Array: ", arr1)

	// initialization of an array with values
	arr2 := [3]int{10, 20, 30}

	fmt.Println("array: ", arr2)

	var slice1 []int

	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)

	fmt.Println("Slice: ", slice1)

	slice2 := []string{"apple", "banana", "orange"}
	fmt.Println("Slice2: ", slice2)

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for index, value := range arr1 {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	k := 0
	for {
		// fmt.Println(k)
		k++
		if k == 5 {
			break
		}
		fmt.Println(k)
	}

	// if else
	x := 10
	y := 20
	if x > 0 {
		if y > 0 {
			fmt.Println("x and y are positive")
		} else {
			fmt.Println("x is positive but y is non-positive")
		}
	} else {
		fmt.Println("x is non-positive")
	}

	score := 75
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else if score >= 50 {
		fmt.Println("F")
	}

	// Switch
	num := 10
	switch num {
	case 10, 20:
		fmt.Println("Number is either 10 or 20")
	case 30, 40:
		fmt.Println("Number is either 30 or 40")
	default:
		fmt.Println("Number is not in the given list")
	}
}
