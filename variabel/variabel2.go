package main
import "fmt"
func var1(){
	// cara 1
	var kalimat = "Kata kan saja!"
	fmt.Println(kalimat)

	// cara 2
	var a1, a2 int = 8, 10
	fmt.Println(a1, a2)

	// cara 3 -> boolean
	var t = true
	fmt.Println(t)

	// cara 4 -> tanpa memberikan nilainya atau zero value
	var e int
	fmt.Println(e)

	// cara 5 -> shorthand
	e1 := "Jeruk"
	fmt.Println(e1)

}