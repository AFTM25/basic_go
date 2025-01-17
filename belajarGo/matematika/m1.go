package main
import "fmt"
func main(){
	a := 16
	b := 20
	c := 11
	fmt.Println("Nilai a = ", a)
	fmt.Println("Nilai b = ", b)
	fmt.Println("Nilai c = ", c)


	hasil := a + b
	fmt.Println("a + b = ",hasil)

	hasil = a - c
	fmt.Println("a - c = ",hasil)

	hasil = b * a
	fmt.Println("b x a = ",hasil)

	hasil = a * b + c
	fmt.Println("a x b + c = ", hasil)

	hasil = a % 5
	fmt.Println("a % 5 = ", hasil)
	fmt.Println()

	// augmented assignment
	i := 10
	fmt.Println(i)
	i += 10
	fmt.Println(i)
	fmt.Println()

	// unary operator
	j := 3
	fmt.Println(j)
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
	
}