package main

import "fmt"

func pangkatt(base, pangkat int) int {
	total := 1
	for i := 1; i <= pangkat; i++ {
		total = total * base
	}
	return total
}

func main() {
	var base int
	var pangkat int

	fmt.Print("Masukan Bilangan (x): ")
	fmt.Scanf("%d\n", &base)
	fmt.Print("Masukan Pangkat (n): ")
	fmt.Scanf("%d", &pangkat)

	fmt.Println("----->>>>><<<<<-----")
	fmt.Print("Hasil Screning : ")
	result := pangkatt(base, pangkat)
	fmt.Println(result)

	//fmt.Println(pangkat(5, 3))
	//fmt.Println(pangkat(10, 2))
	//fmt.Println(pangkat(2, 5))
	//fmt.Println(pangkat(7, 3))
}
