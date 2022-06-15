package main

import "fmt"

func main() {

	// Input
	var bilangan int
	fmt.Println("Masukan Faktor Bilangan")
	fmt.Scanf("%d", &bilangan)

	// Proses
	fmt.Printf("Hasil faktorisasi dari bilangan %d adalah : \n", bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
