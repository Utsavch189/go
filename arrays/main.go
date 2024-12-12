package main

import "fmt"

func main() {
	// Declare an array of size 5 with integer elements
	var arr [5]int
	fmt.Println("Default array:", arr) // [0 0 0 0 0]

	// Using an Array Literal
	arr1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Initialized array:", arr1) // [10 20 30 40 50]

	//  Automatic Size Determination
	arr2 := [...]int{10, 20, 30}                   // Size determined automatically
	fmt.Println("Array with inferred size:", arr2) // [10 20 30]

	// Iterating Over an Array

	//  Using a Standard Loop
	arr3 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("Index %d: %d\n", i, arr3[i])
	}

	//  Using range
	for index, value := range arr3 {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Array Properties

	/*
		Fixed Size: The size of an array cannot be changed.
		Value Type: Arrays are value types in Go. Assigning an array to another creates a copy.
	*/
	a := [3]int{1, 2, 3}
	b := a // Copy of `a` is created
	b[0] = 99
	fmt.Println("Array a:", a) // [1 2 3]
	fmt.Println("Array b:", b) // [99 2 3]

	// Multidimensional Arrays

	// Declaring a 2D Array
	var matrix [2][3]int
	fmt.Println("Default 2D array:", matrix) // [[0 0 0] [0 0 0]]

	// Initializing a 2D Array
	matrix1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D array:", matrix1)
	fmt.Println("Element at [1][2]:", matrix1[1][2]) // 6

	// Finding the Length
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Length of the array:", len(arr4)) // 5

	// Comparing Arrays [Arrays can be compared if they are of the same size and type.]
	a1 := [3]int{1, 2, 3}
	b1 := [3]int{1, 2, 3}
	c := [3]int{4, 5, 6}

	fmt.Println(a1 == b1) // true
	fmt.Println(a1 == c)  // false

}
