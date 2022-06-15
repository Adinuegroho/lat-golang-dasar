package main

import "fmt"

func playWithAsterisk(n int) {
	for x := 1; x <= n; x++ {
		for y := n; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterisk(5)
}
