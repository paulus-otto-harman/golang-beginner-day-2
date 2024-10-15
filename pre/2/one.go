package main

import "fmt"

func main() {
	jam := 14
	switch {
	case jam >= 12 && jam <= 13:
		fmt.Println("makan siang")
	case jam >= 9 && jam < 12 || jam >= 13 && jam <= 15:
		fmt.Println("belajar")
	}

	hari := 1
	switch hari {
	case 1, 2, 3:
		fmt.Println("lembur")
	case 4, 5:
		fmt.Println("tepat waktu")
	case 6, 7:
		fmt.Println("libur")
	}

}
