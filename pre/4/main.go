package main

import "fmt"

func sum(deret []int) int {
	return len(deret)
}

func main() {
	var coba = []int{1, 2, 3}

	fmt.Println(len(coba), cap(coba), coba)

	coba = append(coba, 4, 5)

	fmt.Println(len(coba), cap(coba), coba)

	fmt.Println(len(coba[1:2]), cap(coba[1:2]), coba[1:2])
}
