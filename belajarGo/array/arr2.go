package main

import "fmt"
func arr3(){
	dataNama := [...]string{"Andi", "Rusmana", "Anwar", "Lamekh", "Jaka", "Ridho", "Salsa", "Budi", "Ridwan", "Akbar"}
	sliceData1 := dataNama[2:7]
	fmt.Println("Data nama  : ", dataNama)
	fmt.Println("Data slice : ", sliceData1)

	// mengubah data nama dengan slice
	sliceData1[1] = "Aditya"
	sliceData1[3] = "Yoga"
	fmt.Println("Setelah diubah = ", sliceData1)
	fmt.Println("Setelah diubah = ", dataNama)

	// menambah data nama
	tambahData := append(sliceData1, "Kurniawan")
	fmt.Println("Penambahan = ", tambahData)
	fmt.Println("Data nama  = ", dataNama)
	fmt.Println("Data slice = ", sliceData1)

	// mengecek masing2 kapasitas
	fmt.Println()
	fmt.Println("==== Kapasitas ====")
	fmt.Println("Pada data Nama adalah ", cap(dataNama))
	fmt.Println("Pada slice adalah ", cap(sliceData1))
	fmt.Println("Pada penambahan adalah ", cap(tambahData))

}