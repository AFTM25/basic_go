package main

import "fmt"
func main(){
    // Kode program if statement
    nama := "Andi"
    if nama == "Andi"{
        fmt.Println("Hello Andi")
    }
    
    // kode program if else statement
    nama1 := "Adi"
    if nama1 == "Aditya"{
        fmt.Println("Hello Aditya")
    } else {
        fmt.Println("Kenalan dulu lah..")
    }
    fmt.Println()
    
    // Kode prorgam if else if else statement
    angka := 12
    if angka > 12{
        fmt.Println("Angka lebih dari 12")
    } else if angka == 12{
        fmt.Println("Angka sama dengan 12")
    } else {
        fmt.Println("Angka kurang dari 12")
    }
    
    fmt.Println()
    
    // kode program shorthand if statement
    // menggunakan if biasa
    length := len(nama1)
    if length > 5{
        fmt.Println("Huruf dalam nama terlalu banyak")
    } else{
        fmt.Println("Huruf sudah sesuai")
    }
    
    // menggunakan shorthand
    if angka1 := 15; angka >= angka1{
        fmt.Println("Angka lebih dari 15")
    } else{
        fmt.Println("Angka kurang dari 15")
    }
}