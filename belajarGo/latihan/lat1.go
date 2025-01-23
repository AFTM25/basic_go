package main
import "fmt"
func main(){
	const harga_barang float32 = 20.0
	var harga_input, hasil, diskon float32
	fmt.Printf("Masukkan Harga (dalam satuan) : ")
	fmt.Scan(&harga_input)

	switch{
		case harga_input > harga_barang:
			diskon = harga_input * 10/100
			hasil = harga_input - diskon
			fmt.Printf("Diskon : Rp. %.3f\n",diskon)
			fmt.Printf("Bayar : Rp. %.3f\n", hasil)
		default:
			fmt.Println("Anda tidak mendapatkan diskon")
	}
}