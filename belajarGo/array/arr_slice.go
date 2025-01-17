package main

import "fmt"
func slice_arr1(){
	names := [...]string{"Aditya", "Lukman", "Rara", "Nugraha", "Jaka", "Laras"}
	
	// mengambil elemen array dari indeks ke 4 hingga indeks sebelum 6 (6 -1)
	slice1 := names[4:6]
	fmt.Println(slice1)

	// mengambil elemen array dari indeks pertama hingga indeks sebelum 3(3 - 1)
	slice2 := names[:3]
	fmt.Println(slice2)

	// mengambil elemen array dari indeks ke 3 hingga indeks terakhir
	slice3 := names[3:]
	fmt.Println(slice3)

	// mengambil semua elemen array
	slice4 := names[:]
	fmt.Println(slice4)

	// Pembuatan secara manual nya
	// var slice5 []string = names[:]

	fmt.Println()
	dataHari := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	s1 := dataHari[2:5]
	fmt.Println("Hari[2:5] = ", s1)
	fmt.Println("Jumlah item [2:5] = ", len(s1))
	fmt.Println("Capacity[2:5] = ", cap(s1))

	s2 := dataHari[:4]
	fmt.Println("Hari[:4] = ", s2)
	fmt.Println("Jumlah item[:4] = ", len(s2))
	fmt.Println("Capacity[:4] = ", cap(s2))

	fmt.Println("Jumlah item pada array = ", len(dataHari))
	fmt.Println()

	// perubahan pada data slice sama dengan perubahan pada data array
	sliceHari := dataHari[5:]
	fmt.Println(sliceHari)
	sliceHari[0] = "Sabtu baru"
	sliceHari[1] = "Minggu Baru"
	fmt.Println(sliceHari)
	fmt.Println(dataHari)

	// menambahkan data baru 
	sliceHari1 := append(sliceHari, "Libur baru")
	fmt.Println(sliceHari1)

	
}