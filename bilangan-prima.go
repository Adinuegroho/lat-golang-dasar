package main

import "fmt"

func primeNumber(number int) bool {
	faktor := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			faktor = faktor + 1
		}
	}

	if faktor == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	var input int

	fmt.Print("Masukan Bilangan : ")
	fmt.Scanf("%d", &input)

	fmt.Println("----->>>>><<<<<-----")
	fmt.Print("Hasil Screning : ")
	result := primeNumber(input)
	if result == true {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}

	//fmt.Println(primeNumber(11))
	//fmt.Println(primeNumber(13))
	//fmt.Println(primeNumber(7))
	//fmt.Println(primeNumber(17))
	//fmt.Println(primeNumber(20))
	//fmt.Println(primeNumber(35))
}
