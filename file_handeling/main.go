package main

import (
	"fmt"
	"io"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFile(filename string, content string) int {
	file, err := os.Create(filename)
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	defer file.Close()
	return length
}

func appendToFile(filename string, data string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}

func readFile(filename string) string {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	return string(databyte)
}

func main() {
	// fmt.Println("File writing...")
	// content := "Hello world!"
	// length := writeFile("./utsav.txt", content)
	// fmt.Println("Created file length is : ", length)
	// fmt.Println("File Reading...")
	// data := readFile("./utsav.txt")
	// fmt.Println("File Data is : ", data)
	fmt.Println("Data appending...")
	appendToFile("./utsav.txt", "second line!!!")
	fmt.Println("File Reading...")
	data := readFile("./utsav.txt")
	fmt.Println("File Data is : ", data)
}
