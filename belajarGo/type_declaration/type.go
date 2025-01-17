package main

import "fmt"
func main(){
	// keyword type declaration
	type angka int8
	var angka1 angka = 89
	fmt.Println(angka1)

	type str1 string
	var str2 str1 = "Aditya"
	fmt.Println("Kata ke dua : ", string(str2[2]))
}