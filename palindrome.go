package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		j := len(input) - 1 - i
		if input[i] != input[j] {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Print("Masukan Kata : ")
	fmt.Scanf("%s", &input)

	fmt.Println("----->>>>><<<<<-----")
	fmt.Print("Hasil Screning : ")
	result := palindrome(input)
	if result == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}

	//fmt.Println(palindrome("lion"))
	//fmt.Println(palindrome("civic"))
}
