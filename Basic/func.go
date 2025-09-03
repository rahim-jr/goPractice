package main

import "fmt"

// function structure
// func functionName (function parameters) return type {} -> single return
// func functionName (function parameters) (return type , return type) {} -> multiple return

// Printing Hello message
func printHello() {
	fmt.Println("System Starting...............")
}

// User name
func getName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

// Taking user input
func getTowNumbers() (int, int) {
	var number1 int
	var number2 int

	fmt.Print("Input Number1 to the system -> ")
	fmt.Scanln(&number1)

	fmt.Print("Input Number2 to the system -> ")
	fmt.Scanln(&number2)

	return number1, number2
}

func greeting(user string) {
	fmt.Println("Hello " + user + ", Have a nice day!")
}

// Helping the user
func add(a int, b int) {
	fmt.Println("Sum of the two numbers is: ", a+b)
}

func goodByeUser(name string) {
	fmt.Println("Good Bye! " + name)
}

func init() {
	printHello()
}

func main() {
	userName := getName()
	number1, number2 := getTowNumbers()
	greeting(userName)
	add(number1, number2)
	goodByeUser(userName)
}
