package main

import "fmt"
func func_slice(){
	// Membuat slice dengan make
	newSlice := make([]int, 2, 5)
	newSlice[0] = 12
	newSlice[1] = 15
	// newSlice[2] error

	newSlice2 := append(newSlice, 89)
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
}