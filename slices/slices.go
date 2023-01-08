package slices

import "fmt"

func Pizza() {
	isimler := make([]string, 3)

	fmt.Println(isimler)

	isimler[0] = "Hülya"
	isimler[1] = "Almila"
	isimler[2] = "Mustafa"
	// isimler[4] = "Gökalp"    bu satır eklenirse  panic: runtime error: index out of range [3] with length 3 hatası verir
	isimler = append(isimler, "KAYA")

	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
