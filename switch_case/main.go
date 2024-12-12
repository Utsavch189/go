package main

import "fmt"

/*
switch variable {
case value1:
    // Code to execute if variable == value1
case value2:
    // Code to execute if variable == value2
default:
    // Code to execute if none of the above cases match
}
*/

func main() {
	// Simple switch Example
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the workweek.")
	case "Friday":
		fmt.Println("Almost the weekend!")
	case "Saturday", "Sunday": // Multiple cases
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a regular weekday.")
	}

	// Using switch Without a Variable [You can write switch without specifying a variable, which allows you to use complex conditions.]
	age := 20

	switch {
	case age < 18:
		fmt.Println("You are a minor.")
	case age >= 18 && age < 60:
		fmt.Println("You are an adult.")
	case age >= 60:
		fmt.Println("You are a senior citizen.")
	}

	// fallthrough Keyword [In Go, cases do not fall through by default. To explicitly allow fallthrough to the next case, use the fallthrough keyword.]
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("Excellent!")
		fallthrough
	case "B":
		fmt.Println("Well done!")
		fallthrough
	case "C":
		fmt.Println("You passed!")
	default:
		fmt.Println("Try harder next time.")
	}
	/*
		OUTPUT:
			Well done!
			You passed!
			Try harder next time.
	*/

	// Switch with Multiple Values [You can match multiple values for a single case.]
	color := "blue"

	switch color {
	case "red", "blue", "green":
		fmt.Println("You chose a primary color.")
	default:
		fmt.Println("You chose a different color.")
	}

	// Switch with Type Assertions [Go's switch can handle types in interface{} values using a type switch.]
	var value interface{} = "42"

	switch v := value.(type) {
	case int:
		fmt.Printf("Value is an int: %d\n", v)
	case string:
		fmt.Printf("Value is a string: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}

	/*
		Key Points:
			No Fallthrough by Default: Go's switch does not fall through unless explicitly told with the fallthrough keyword.
			Default is Optional: Use it for handling unmatched cases.
			Multiple Cases: You can combine cases for the same outcome.
			No Break Required: Each case block ends automatically.
			Type Switch: Useful for handling dynamic types in interface{}.
	*/
}
