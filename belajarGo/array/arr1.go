package main

import "fmt"
func coba_arr1(){
	values := [...]int{3,2,1}
	fmt.Println(values)
	fmt.Println("Jumlah array : ", len(values))
	
	values[0] = 32
	fmt.Println(values)
	fmt.Println("Jumlah array : ", len(values))

}