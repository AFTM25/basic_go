package main

import "fmt"
func main(){
	type nama string

	// deklarasi array
	var name[3] nama

	// mengisi array
	name[0] = "Aditya"
	name[1] = "Rudi"
	name[2] = "Anwar"

	// mengakses array
	fmt.Println("Nama 1 : ", name[0])
	fmt.Println("Nama 2 : ", name[1])
	fmt.Println("Nama 3 : ", name[2])

	fmt.Println()

	var values = [3]string{
		"Aku",
		"Kamu",
		"Dia",
	}
	fmt.Println(values)

	val1 := [2]string{"aku", "saya"}
	fmt.Println(val1)
	fmt.Println()

	// Mengakses function dari file arr1.go
	coba_arr1()
	fmt.Println()

	// Mengakses function dari file arr_slice.go
	slice_arr1()

	// Mengakses function pada satu file
	kamu()

	arr3()

	fmt.Println()
	func_slice()
}

// Membuat func baru dalam satu file
func kamu(){
	fmt.Println()
	fmt.Println("Hello!!!!!")
	fmt.Println()
}