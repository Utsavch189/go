package main

import "fmt"

func main() {
	// Declaring a Map
	// Using make()
	m := make(map[string]int) // Map with string keys and int values
	fmt.Println(m)            // Output: map[]

	// Using a Map Literal
	m1 := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println(m1) // Output: map[Alice:25 Bob:30]

	// Adding or Updating Elements
	m2 := make(map[string]int)
	m2["Alice"] = 25 // Add a new key-value pair
	m2["Bob"] = 30   // Add another key-value pair
	m2["Alice"] = 26 // Update the value for the key "Alice"
	fmt.Println(m2)  // Output: map[Alice:26 Bob:30]

	// Retrieving Values
	age := m2["Alice"]
	fmt.Println(age) // Output: 26

	// Checking if a Key Exists
	value, ok := m2["Charlie"]
	if ok {
		fmt.Println("Key exists with value:", value)
	} else {
		fmt.Println("Key does not exist")
	}

	// Deleting a Key
	delete(m2, "Alice") // Remove the key "Alice"
	fmt.Println(m2)     // Output: map[Bob:30]

	// Iterating Over a Map
	m3 := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	for key, value := range m3 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Nil Maps [A nil map is a map that has not been initialized. Attempting to add elements to a nil map will cause a runtime panic]
	var m4 map[string]int // Declare the map, but it is nil
	fmt.Println(m4)       // Output: map[]
	m4["Alice"] = 25      // Panic: assignment to entry in nil map
	// So to fix the panic
	m4 = make(map[string]int) // Initialize the map
	m4["Alice"] = 25          // Now this works
	fmt.Println(m4)           // Output: map[Alice:25]

	// Empty Maps [An empty map is initialized but has no elements:]
	m5 := make(map[string]int) // Empty map
	fmt.Println(len(m5))       // Output: 0

}
