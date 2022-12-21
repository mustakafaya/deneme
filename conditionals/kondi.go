package kondip

import "fmt"

func Kondif() {

	bakiye := 2500.25
	cek := 1600.25

	if bakiye < cek {

		fmt.Println("para yoh")

	} else if bakiye == cek {

		bakiye = bakiye - cek

		fmt.Println("Paranız hazırlanıyor kalan bakiyeniz:", bakiye, "₺")
		// fmt.Println("paranız hazırlanıyor kalan bakiye : " + fmt.Sprintf("%v", bakiye))
	} else {

		bakiye = bakiye - cek
		fmt.Println("paranız hazırlanıyor kalan bakiye : " + fmt.Sprintf("%v", bakiye))

	}
}
