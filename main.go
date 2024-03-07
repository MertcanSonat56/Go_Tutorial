package main

import "fmt"

// Define a struct type representing a person
type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street     string
	City       string
	PostalCode string
}

// Function to calculate the sum of two integers
func add(a, b int) int {
	return a + b
}

// Function to calculate the difference of two integers
func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	return float64(a) / float64(b), nil
}

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

	// functions
	num1 := 10
	num2 := 5

	sum := add(num1, num2)
	fmt.Println(sum)

	difference := subtract(num1, num2)
	fmt.Println(difference)

	product := multiply(num1, num2)
	fmt.Println(product)

	quotient, err := divide(num1, num2)

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(quotient)
	}

	// Maps

	// initializing a map with key-value pairs
	departmentHeadcount := map[string]int{
		"Engineering": 10,
		"Marketing":   8,
		"Sales":       12,
	}

	fmt.Println("Engineering: ", departmentHeadcount["Engineering"])
	fmt.Println("Marketing: ", departmentHeadcount["Marketing"])
	fmt.Println("Sales: ", departmentHeadcount["Sales"])

	// Checking if a key exists in the map
	_, ok := departmentHeadcount["HR"]
	if ok {
		fmt.Println("\nHR Department Exists")
	} else {
		fmt.Println("\nHR Department Does Not Exists")
	}

	// Deleting a key-value pair from the map
	delete(departmentHeadcount, "Marketing")
	fmt.Println("\n Department Headcount after deletion: ")
	fmt.Println(departmentHeadcount)

	// Structs
	person1 := Person{
		Name: "Wawex",
		Age:  22,
		Address: Address{
			Street:     "street 123",
			City:       "İstanbul",
			PostalCode: "3434",
		},
	}

	fmt.Println("Name: ", person1.Name)
	fmt.Println("Age: ", person1.Age)
	fmt.Println("Address: ", person1.Address.Street)
	fmt.Println("City: ", person1.Address.City)
	fmt.Println("Postal Code: ", person1.Address.PostalCode)

	// update fields of the person
	person1.Name = "Mertcan Yoldas"
	person1.Age = 35
	person1.Address.City = "İzmir"

	// print updated fields
	fmt.Println("\nUpdated Name: ", person1.Name)
	fmt.Println("\nUpdated Age: ", person1.Age)
	fmt.Println("\nUpdated City: ", person1.Address.City)
}
