package main
import "fmt"
func main(){
	var nilai32 int32 = 32768

	// number overflow, melebihi kapasitas, output nya jadi negatif
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println()

	// konversi string
	name := "Aditya"
	fmt.Println(name)
	
	e := name[0]
	fmt.Println(e)

	eString := string(e)
	fmt.Println(eString)
}