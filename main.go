package main

import (
	"fmt"
	"moduleadi/functions"
)

// "moduleadi/slices"

func main() {

	say1 := 1
	say2 := 1
	fmt.Print("1. sayıyı giriniz;")
	fmt.Scanln(&say1)
	fmt.Print("2. sayıyı giriniz;")
	fmt.Scanln(&say2)

	sonuc1, sonuc2, sonuc3, sonuc4 := functions.Fonks(say1, say2)

	fmt.Println("Toplamı : ", sonuc1)
	fmt.Println("Farkı : ", sonuc2)
	fmt.Println("Çarpımı : ", sonuc3)
	fmt.Println("Bölümü : ", sonuc4)

}
