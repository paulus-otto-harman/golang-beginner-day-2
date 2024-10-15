package main

import "fmt"

func prima(deret []int) []int {
	bilanganPrima := []int{}

	for _, angka := range deret {
		for i := 2; i < angka/2; i++ {
			if angka%i == 0 {
				fmt.Println(angka)
			}
		}
	}

	return bilanganPrima

}

func main() {
	bilangan := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	//fmt.Println(prima(bilangan))
}
