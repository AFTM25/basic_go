package main

import "fmt"
func main(){
	// cara 1
	var name string
	var usia int8
	name = "Aditya"
	usia = 20
	fmt.Println("Nama saya adalah ", name)
	fmt.Println("Usia saya adalah ", usia, " tahun")
	fmt.Println()

	// cara 2 saat inisialisai variabel tipe data
	var kalimat = "Hello Variabel Golang"
	fmt.Println(kalimat)

	kalimat = "Ini adalah pengubahan"
	fmt.Println(kalimat)

	var angka = 90
	fmt.Println(angka)
	fmt.Println()

	// Cara 3, tidak menggunakan keyword var
	nama := "Aditya Pratama Yoga"
	fmt.Println("Nama : ", nama)

	nama = "Golang"
	fmt.Println("User : ", nama)

	/* Note :
	kalau sudah deklarasikan variabel nya, maka untuk pengubahan variabel tidak perlu menuliskan := lagi.
	*/
	fmt.Println()

	// Multiple Variabel
	var (
		first_name = "Aditya Pratama"
		last_name = "Yoga"
		usia1 = 20
		angka_float = 34.90
	)
	fmt.Println(first_name, last_name)
	fmt.Println("Usia = ", usia1, " tahun")
	fmt.Println("Angka Float = ", angka_float)
	fmt.Println()
	var1()
}