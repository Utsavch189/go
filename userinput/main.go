package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Your name : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // comma ok || err syntax. I go there no try except block every error or success treat as a variable.
	fmt.Println("Your name is : ", input)
	fmt.Println("Error : ", err)
}
