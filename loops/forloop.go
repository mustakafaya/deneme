package loops

import "fmt"

func Dongu() {

	count := 77
	tahmin := 0
	kere := 1
	fmt.Println("1 ile 100 arasındaki tahmininizi giriniz")
	fmt.Scanln(&tahmin)

	for tahmin != count {

		if tahmin > count {
			fmt.Println("Daha küçük bir sayı giriniz")
		} else if tahmin < count {
			fmt.Println("Daha büyük bir sayı giriniz")
		}

		fmt.Scanln(&tahmin)

		kere = kere + 1

	}
	if tahmin == count {
		fmt.Print("tebrikler ", kere, " tahminde bildiniz")
		if kere < 4 {

			fmt.Println(": Süper")

		} else {
			fmt.Println(": Fena değil")
		}
	}
}
