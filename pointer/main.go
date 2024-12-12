package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Initial pointer ptr value is : ", ptr) // <nil>

	number := 5
	var ptr2 = &number                      //pointer ptr2 referencing to number variable
	fmt.Println("ptr2 value is : ", ptr2)   // returns a memory address
	fmt.Println("*ptr2 value is : ", *ptr2) // returns actual value

	*ptr2 = *ptr2 * 3 // that will actually change the value of number also
	fmt.Println("value of *ptr2 : ", *ptr2)
	fmt.Println("value of number : ", number)
}
