package main

import "fmt"

/*
Syntax:
func functionName(parameters) returnType {
    // function body
    return value
}
*/

// A simple function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Parameters
// 1. Single Param
func square(n int) int {
	return n * n
}

// 2. Multiple Param
func add(a int, b int) int {
	return a + b
}

// 3. Same Type Parameters [If multiple parameters have the same type, you can declare it concisely.]
func multiply(a, b int) int {
	return a * b
}

// Return Values
// 1.Single Return Value:
func subtract(a, b int) int {
	return a - b
}

// 2.Multiple Return Values:
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Named Return Values [You can name the return variables and omit the return arguments]
func rectangleProperties(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

// Variadic Functions [Variadic functions accept a variable number of arguments]
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Closures [A closure is a function that captures variables from its surrounding scope.]
func makeMultiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

// Panic and Recover [Functions can handle unexpected conditions using panic and recover]
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("Division by zero!")
	}
	fmt.Println("Result:", a/b)
}

// Higher-Order Functions [Functions can take other functions as parameters or return functions.]
// 1.Passing a Function:
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 2.Returning a Function:
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Function Pointers [In Go, functions are first-class citizens and can be assigned to variables or passed as arguments]
func hello() {
	fmt.Println("Hello, Go!")
}

func main() {
	// A simple function
	message := greet("Alice")
	fmt.Println(message)

	// Multiple Return Values:
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// Named Return Values
	a, p := rectangleProperties(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)

	// // Variadic Functions
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", result)

	// Anonymous Functions [Anonymous functions are declared without a name and can be used immediately or assigned to a variable.]
	// 1.Immediate Invocation:
	res := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println("Result:", res)
	// 2.Assigned to a Variable:
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum:", add(5, 7))

	// Closure
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("Double of 4:", double(4)) // 8
	fmt.Println("Triple of 4:", triple(4)) // 12

	// Defer [The defer statement delays the execution of a function until the surrounding function completes.]
	// Use case: Cleaning up resources (e.g., closing files).
	defer fmt.Println("This will print last")
	fmt.Println("This will print first")

	// Panic and Recover
	safeDivide(10, 2)
	safeDivide(10, 0)

	// Higher Order Functions
	// 1.Passing a Function:
	add1 := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	fmt.Println("Add:", applyOperation(3, 4, add1))
	fmt.Println("Multiply:", applyOperation(3, 4, multiply))

	// 2.Returning a Function:
	addFive := makeAdder(5)
	fmt.Println("Add 5 to 3:", addFive(3)) // 8

	// Function Pointers
	var f func()
	f = hello
	f()
}

/*
Summary:
	Functions in Go are versatile, supporting named returns, variadic parameters, and closures.
	They form the basis for higher-order functions and functional programming patterns.
	Use defer, panic, and recover for resource management and error handling.
*/
