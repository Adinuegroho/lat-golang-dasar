package main

import "fmt"

func main() {
	// Deklarasi Variable
	var studentName string
	var studentScore int

	// Input
	fmt.Scanf("%s\n", &studentName)
	fmt.Scanf("%d", &studentScore)

	// Proses
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Printf("Nama : %s, Nilai : %d ,Mendapatkan Grade A", studentName, studentScore)
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Printf("Nama : %s, Nilai : %d ,Mendapatkan Grade B", studentName, studentScore)
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Printf("Nama : %s, Nilai : %d ,Mendapatkan Grade C", studentName, studentScore)
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Printf("Nama : %s, Nilai : %d ,Mendapatkan Grade D", studentName, studentScore)
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Printf("Nama : %s, Nilai : %d ,Mendapatkan Grade E", studentName, studentScore)
	} else {
		fmt.Println("Nilai Invalid")
	}
}
