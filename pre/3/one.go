package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	fmt.Println(add(1, 3))
}
