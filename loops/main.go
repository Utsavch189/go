package main

import "fmt"

func main() {

	// Simple Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// for as a While Loop [If you omit the initialization and post sections, for works like a while loop]
	count := 1

	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

	//  Infinite Loop [A loop with no condition becomes an infinite loop.]
	count1 := 0

	for {
		fmt.Println("Infinite Loop, count:", count1)
		count1++
		if count1 == 5 {
			break // Exiting the loop
		}
	}

	// Iterating Over Collections [Use range to iterate over elements in arrays or slices.]
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterating Over a Map [range can also be used to iterate over key-value pairs in a map]
	data := map[string]int{"A": 1, "B": 2, "C": 3}

	for key, value := range data {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Iterating Over a String
	str := "hello"

	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// Breaking and Continuing Loops
	// 1.break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	// 2.continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Nested Loops [Loops can be nested, and you can use labels to break out of a specific loop.]
	// EX-1:
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break outerLoop
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// EX-2:
outerLoop1:
	for i := 1; i <= 3; i++ {
	outerLoop2:
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				if i == 2 && j == 2 && k == 2 {
					fmt.Println("Breaking out of outerLoop2")
					fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
					break outerLoop2
				} else {
					fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
					break outerLoop1
				}

			}
		}
	}
	fmt.Println("All loops completed")
}

/*
Best Practices:
	Avoid Infinite Loops: Always ensure a termination condition for your loops.
	Use range for Collections: Makes your code cleaner and easier to understand.
	Labels for Clarity: Use labeled loops sparingly for better readability.
	Optimize Loops: Avoid unnecessary computations inside the loop body.
*/
