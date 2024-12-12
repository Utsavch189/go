package main

import (
	"encoding/json"
	"fmt"
)

// Defining a Struct
type Person struct {
	Name   string
	Age    int
	Gender string
}

// Methods on Structs

// Method with a Value Receiver
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s.\n", p.Name)
}

// Method with a Pointer Receiver
func (p *Person) IncrementAge() {
	p.Age++
}

func main() {
	//Creating and Initializing Structs

	// 1. Using a Literal
	p := Person{
		Name:   "Alice",
		Age:    25,
		Gender: "Female",
	}
	fmt.Println(p) // Output: {Alice 25 Female}

	// 2. Named Fields [You can specify only some fields, and the rest will have zero values:]
	p1 := Person{Name: "Bob"}
	fmt.Println(p1) // Output: {Bob 0 }

	// 3. Zero Value
	var p2 Person
	fmt.Println(p2) // Output: { 0 }

	// 4. Using new [The new keyword returns a pointer to a newly allocated zero value struct.]
	p3 := new(Person)
	p3.Name = "Charlie"
	fmt.Println(p3)  // Output: &{Charlie 0 }
	fmt.Println(*p3) // Output: {Charlie 0 }

	// Accessing and Modifying Fields
	p4 := Person{Name: "Alice", Age: 25}
	fmt.Println(p4.Name) // Output: Alice

	p4.Age = 26
	fmt.Println(p4.Age) // Output: 26

	// Struct with Pointers [You can use a pointer to a struct to modify its fields directly.]
	p5 := &Person{Name: "David", Age: 30}
	p5.Age = 31     // Modifies the original struct
	fmt.Println(p5) // Output: &{David 31 }

	// Anonymous Structs [You can define and use structs without giving them a name]
	p6 := struct {
		Name string
		Age  int
	}{
		Name: "Eve",
		Age:  28,
	}
	fmt.Println(p6) // Output: {Eve 28}

	// Structs with Embedded Fields
	type Address struct {
		City  string
		State string
	}

	type User struct {
		Name    string
		Age     int
		Address // Embedded struct
	}

	p7 := User{
		Name: "Frank",
		Age:  40,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}
	fmt.Println(p7) // Output: {Frank 40 {New York NY}}

	// Methods on Structs

	// Method with a Value Receiver [You can define methods on structs by associating them with a receiver.]
	p8 := Person{Name: "Alice"}
	p8.Greet() // Output: Hello, my name is Alice.

	// Method with a Pointer Receiver [Use a pointer receiver when you need to modify the struct.]
	p9 := Person{Name: "Bob", Age: 25}
	p9.IncrementAge()
	fmt.Println(p9.Age) // Output: 26

	// Structs in Slices and Maps

	// Slice of Structs
	people := []Person{
		{Name: "Alice", Age: 25, Gender: "Female"},
		{Name: "Bob", Age: 30, Gender: "Male"},
	}
	fmt.Println(people) // Output: [{Alice 25 Female} {Bob 30 Male}]

	// Map of Structs
	people1 := map[string]Person{
		"Alice": {Name: "Alice", Age: 25},
		"Bob":   {Name: "Bob", Age: 30},
	}
	fmt.Println(people1["Alice"]) // Output: {Alice 25}

	// Tagging Struct Fields [Struct fields can have tags, which are metadata used by libraries like encoding/json]
	type Person1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p10 := Person1{Name: "Alice", Age: 25}
	b, _ := json.Marshal(p10)
	fmt.Println(string(b)) // Output: {"name":"Alice","age":25}

	/*
		Best Practices:
			Use pointer receivers for methods if the struct is large or needs modifications.
			Embed structs for reusability and composition.
			Use tags to customize struct behavior in libraries like JSON, XML, and database drivers.
	*/
}
