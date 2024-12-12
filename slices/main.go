package main

import "fmt"

/*
What is a Slice?
A slice is a lightweight, dynamic abstraction over an array. Unlike arrays, slices do not have a fixed size. They describe a contiguous section of an underlying array and include:

A pointer to the array.
A length (number of elements in the slice).
A capacity (maximum number of elements before resizing).
*/

func main() {
	// Declaring a Slice

	// 1. Using a Slice Literal
	slice := []int{1, 2, 3} // Creates a slice with 3 elements
	fmt.Println(slice)      // Output: [1 2 3]

	// 2. From an Array
	arr := [5]int{10, 20, 30, 40, 50}
	slice1 := arr[1:4]  // Slice from index 1 to 3 (excludes index 4)
	fmt.Println(slice1) // Output: [20 30 40]

	// 3. Using make() [You can create a slice with a specific length and capacity using make()]
	slice2 := make([]int, 5, 10) // Length 5, Capacity 10
	fmt.Println(slice2)          // Output: [0 0 0 0 0]
	fmt.Println(len(slice2))     // 5
	fmt.Println(cap(slice2))     // 10

	// Accessing and Modifying Slice Elements [You can access and modify elements of a slice just like arrays:]
	slice3 := []int{10, 20, 30}
	fmt.Println(slice3[1]) // Output: 20
	slice3[1] = 25
	fmt.Println(slice3) // Output: [10 25 30]

	// Appending to a Slice
	slice4 := []int{1, 2, 3}
	slice4 = append(slice4, 4) // Add one element
	fmt.Println(slice4)        // Output: [1 2 3 4]

	slice4 = append(slice4, 5, 6, 7) // Add multiple elements
	fmt.Println(slice4)              // Output: [1 2 3 4 5 6 7]

	// Appending One Slice to Another
	slice5 := []int{1, 2, 3}
	slice6 := []int{4, 5, 6}
	slice1 = append(slice5, slice6...) // Use `...` to unpack elements
	fmt.Println(slice5)                // Output: [1 2 3 4 5 6]

	// Slicing a Slice [You can create a new slice from an existing slice using the slicing syntax:]
	slice7 := []int{10, 20, 30, 40, 50}
	newSlice := slice7[1:4] // Includes elements from index 1 to 3
	fmt.Println(newSlice)   // Output: [20 30 40]

	/*
		Important Points:
			Slicing a slice creates a new slice that refers to the same underlying array.
			Modifying the new slice will also affect the original slice if they overlap.
	*/
	slice8 := []int{10, 20, 30, 40, 50}
	newSlice1 := slice8[1:4]
	newSlice1[0] = 100
	fmt.Println(slice8)    // Output: [10 100 30 40 50]
	fmt.Println(newSlice1) // Output: [100 30 40]

	/*
		Slice Length and Capacity
			Length (len): Number of elements in the slice.
			Capacity (cap): Number of elements in the underlying array starting from the first element of the slice.
	*/
	slice9 := []int{10, 20, 30, 40, 50}
	newSlice2 := slice9[1:4]    // [20 30 40]
	fmt.Println(len(newSlice2)) // 3 (length of the slice)
	fmt.Println(cap(newSlice2)) // 4 (elements from index 1 to end of array)

	// Resizing and Capacity Growth [When a slice exceeds its capacity, Go creates a new underlying array (with increased size, usually doubled) and copies the data.]
	slice10 := []int{1, 2, 3}
	fmt.Println(cap(slice10))  // Output: 3
	slice = append(slice10, 4) // Adds an element and creates a new array
	fmt.Println(cap(slice10))  // Output: 6 (capacity doubled)

	// Iterating Over a Slice

	// Using for Loop
	slice11 := []int{10, 20, 30}
	for i := 0; i < len(slice11); i++ {
		fmt.Println(slice11[i])
	}

	// Using range
	slice12 := []int{10, 20, 30}
	for index, value := range slice12 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Copying a Slice
	source := []int{10, 20, 30}
	destination := make([]int, len(source))
	copy(destination, source) // Copy elements
	fmt.Println(destination)  // Output: [10 20 30]

	// Multidimensional Slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix[1][2]) // Output: 6

}
