package main

import "fmt"

func cetakTablePerkalian(number int) {
	for x := 1; x <= number; x++ {
		for y := 1; y <= number; y++ {
			fmt.Print(y * x)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func main() {
	cetakTablePerkalian(3)
}
