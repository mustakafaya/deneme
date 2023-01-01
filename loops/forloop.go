package loops

import "fmt"

func Dongu() {

	count := 77
	tahmin := 0
	fmt.Println("1 ile 100 arasındaki tahmininizi giriniz")
	fmt.Scanln(&tahmin)

	for tahmin != count {

		if tahmin > count {
			fmt.Println("Daha küçük bir sayı giriniz")
		} else if tahmin < count {
			fmt.Println("Daha büyük bir sayı giriniz")
		}

		fmt.Scanln(&tahmin)

	}
	if tahmin == count {
		fmt.Println("tebrikelr bildiniz")
	}
}
