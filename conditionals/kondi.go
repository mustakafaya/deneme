package kondip

import "fmt"

func Kondif() {

	bakiye := 2500.25
	cek := 1600.25

	if bakiye < cek {

		fmt.Println("para yoh")

	}

	if bakiye >= cek {

		bakiye = bakiye - cek

		fmt.Println("paranız hazırlanıyor kalan bakiye : ", bakiye)
	}
}
