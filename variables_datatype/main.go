package main

import "fmt"

// global variable declaration needs var keyword
var globalvar int = 789

// OR
var globalvar1 = "hello"

// Constants
const MyConstant = "I am constant!" // first character capital means its a public variable

func main() {
	var name string = "utsav"
	fmt.Println(name)
	fmt.Printf("Type : %T \n", name)

	var isExists bool = true
	fmt.Println(isExists)
	fmt.Printf("Type : %T \n", isExists)

	var signedNumber int = 567858787676 // can be pos or neg
	fmt.Println(signedNumber)
	fmt.Printf("Type : %T \n", signedNumber)

	var unsignedNumber uint = 100 // only positive
	fmt.Println(unsignedNumber)
	fmt.Printf("Type : %T \n", unsignedNumber)

	var smallDecimals float32 = 54.44
	fmt.Println(smallDecimals)
	fmt.Printf("Type : %T \n", smallDecimals)

	var bigDecimals float64 = 54.448778
	fmt.Println(bigDecimals)
	fmt.Printf("Type : %T \n", bigDecimals)

	var smallPosInt uint8 = 255 // 0-255 // Same there has uint16,uint32,uint64
	fmt.Println(smallPosInt)
	fmt.Printf("Type : %T \n", smallPosInt)

	var smallSignedInt int8 = 127 // -127 to 127 // Same there has int16,int32,int64
	fmt.Println(smallSignedInt)
	fmt.Printf("Type : %T \n", smallSignedInt)

	var complex complex128 = 1 + 2i
	fmt.Println(complex)
	fmt.Printf("Type : %T \n", complex)

	// Another type of variable declare [without using type]
	var myname = "utsav chatterjee"
	fmt.Println(myname)
	fmt.Printf("Type : %T \n", myname) // string [without type go can do dynamic typing]

	// Another type of variable declare [without using var]
	username := "utsav1234"
	fmt.Println(username)
	fmt.Printf("Type : %T \n", myname)

	// DERIVED TYPE :

	// Array: A fixed-size sequence of elements of the same type.
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Slice: A dynamic-sized, flexible view into an array
	var scores []int = []int{90, 85, 80}
	fmt.Println(scores)

	//Map: A collection of key-value pairs.
	var person map[string]string = map[string]string{
		"name": "Alice",
		"age":  "25",
	}
	fmt.Println(person)

	//Struct: A custom data structure.
	type Person struct {
		Name string
		Age  int
	}
	var user = Person{"Alice", 25}
	fmt.Println(user)
}
