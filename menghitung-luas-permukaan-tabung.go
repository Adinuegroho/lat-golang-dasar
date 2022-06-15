package main

import "fmt"

func main() {
	// Deklarasi Variable
	var T, r float64
	const phi = 3.14

	// Input
	fmt.Println("Masukan Tinggi Tabung")
	fmt.Scanf("%g\n", &T)
	fmt.Println("Masukan Jari-jari Tabung")
	fmt.Scanf("%g", &r)

	// Proses
	fmt.Println("Hasil Luas Permukaan Tabung")
	var lp = 2 * phi * r * (r + T)
	fmt.Println(lp)
}
