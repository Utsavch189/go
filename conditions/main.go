package main

import "fmt"

func main() {
	// Simple if-else if-else
	marks := 85

	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else if marks >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Declaring Variables Inside if [You can declare and initialize variables directly in the if statement. These variables are scoped only within the if and else blocks.]
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// Nested if Statements
	age := 20
	isCitizen := true

	if age >= 18 {
		if isCitizen {
			fmt.Println("You are eligible to vote.")
		} else {
			fmt.Println("You must be a citizen to vote.")
		}
	} else {
		fmt.Println("You are not old enough to vote.")
	}

	// Using Logical Operators in if [You can combine conditions using logical operators && (AND), || (OR), and ! (NOT)]
	age1 := 25
	income := 50000

	if age1 >= 18 && income > 30000 {
		fmt.Println("You qualify for the loan.")
	} else {
		fmt.Println("You do not qualify for the loan.")
	}
}
