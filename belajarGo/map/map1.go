package main
import "fmt"
func main(){
	fmt.Println("Tipe Data Map")
	/*
	// cara 1
	var person map[string]string = map[string]string{}
	//isi data nya
	person["name"] = "aditya"
	fmt.Println(person)
	*/

    // membua map dengan cara biasa
	dataBuah := map[string]string{
		"buah1" : "Mangga",
		"buah2" : "Apel",
		"buah3" : "Jeruk",
	}

	// mengakses key pada map
	fmt.Println(dataBuah["buah1"])
	fmt.Println(dataBuah["buah2"])
	fmt.Println(dataBuah)
	fmt.Println()
	
	// membuat map dengan fungsi make(map[typeKey]typeValue)
	dataHuruf := make(map[string]string)
	
	// menambahkan data pada map make
	dataHuruf["huruf1"] = "A"
	dataHuruf["huruf2"] = "B"
	dataHuruf["huruf3"] = "C"
	fmt.Println(dataHuruf)
	
	// Mengubah value dengan key
	dataHuruf["huruf1"] = "F"
	fmt.Println(dataHuruf)
	
	// mengetahui jumlah data pada map
	jumlah := len(dataHuruf)
	fmt.Println(jumlah)
	
	// Menghapus data pada map dengan key
	delete(dataHuruf, "huruf3")
	fmt.Println(dataHuruf)
}