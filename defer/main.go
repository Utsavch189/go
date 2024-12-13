package main

import "fmt"

func main() {
	// Defer lines actually goes into a stack and executes last for that scope
	defer fmt.Println("I am defer")
	fmt.Println("Hello")
	/*
		OUTPUT :
			Hello
			I am defer
	*/

	// multiple defers
	defer fmt.Println("I am defer1")
	defer fmt.Println("I am defer2")
	defer fmt.Println("I am defer3")
	fmt.Println("Hello2")
	/*
		 At first defers are going into the stack :
		 So, I am defer1,I am defer2,I am defer3
		 Now Hello2 will be printed
		 after that at very last according to stack's last in first out nature,
		 I am defer3,
		 I am defer2,
		 I am defer1
		 will be printed

		 OUTPUT:
		 	Hello2
			I am defer3
			I am defer2
			I am defer1
	*/

	// Defer for function scope
	defer fmt.Println("I am defer4")
	fmt.Println("Hello3")
	myDefer()
	fmt.Println()
	/*
		At first I am defer4 will go into stack for main func scope
		Then Hello3 will be printed
		While calling myDefer(), its inner defer will be executed into its scope,
		so into that func 0,1,2,3,4 will go into stack
		and accroding stack's lifo nature 43210 will be printed and control will be back into main func.
		At last I am defer4 which was prev stacked into main func's scope will be printed.

		OUTPUT:
			Hello3
			43210
			I am defer4
	*/
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
