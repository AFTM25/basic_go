package main

import "fmt"
func main(){
	// break
	for i := 1; i < 12 ; i++{
		if i == 8{
			break
		}
		fmt.Printf("Perulangan ke %v\n", i)
	}

	fmt.Println()

	// continue
	for i := 0; i < 10; i++{
		if i % 2 == 0{
			continue
		}
		fmt.Printf("Angka : %v\n", i)
	}
}