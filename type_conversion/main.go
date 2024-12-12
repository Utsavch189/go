package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a rating between 1 to 10 : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Before Conversion : ", input)
	fmt.Printf("Type : %T \n", input)
	// Trim whitespace from the input and parse it as a base-10 integer with 64-bit size
	numberFormat, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("After Conversion : ", numberFormat)
		fmt.Printf("Type : %T \n", numberFormat)
	}

	/*
		Primitive types: Use explicit conversion syntax like float64(x).
		String and Numbers: Use strconv for safe conversions.
		Custom types: Conversion is explicit even if the underlying type is the same.
		Go enforces strict type safety, so all conversions must be intentional and explicit.
	*/

	//1. Primitive types :
	var x int = 10
	var y float64 = float64(x)
	fmt.Println(y) // Output: 10.0

	var x1 float64 = 5.67
	var y1 int = int(x1)
	fmt.Println(y1) // Output: 5 (truncates the decimal part)

	//2. String and Numbers :

	// String to Integer (using strconv):
	var str string = "42"
	num, err := strconv.Atoi(str) // Atoi converts string to int
	if err == nil {
		fmt.Println(num) // Output: 42
	}
	/*
		Use strconv.Atoi for simple decimal integer parsing when int type suffices.
		Use strconv.ParseInt for advanced parsing needs, such as different bases or explicit control over bit sizes.
	*/

	// Integer to String (using strconv):
	var num1 int = 42
	var str1 string = strconv.Itoa(num1) // Itoa converts int to string
	fmt.Println(str1)                    // Output: "42"

	//3. Custom types:

	// Byte Slice to String:
	var data = []byte{'H', 'e', 'l', 'l', 'o'}
	str2 := string(data)
	fmt.Println(str2) // Output: "Hello"

	// String to Byte Slice:
	var str3 = "Hello"
	bytes := []byte(str3)
	fmt.Println(bytes) // Output: [72 101 108 108 111]

	//4. Custom Types
	// When working with custom types, Go enforces explicit conversion even if the underlying type is the same.
	type MyInt int
	var x2 int = 10
	var y2 MyInt = MyInt(x2) // Explicit conversion required
	fmt.Println(y2)          // Output: 10

}
