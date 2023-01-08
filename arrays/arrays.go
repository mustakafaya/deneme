package arrays

import "fmt"

func Dizi() {

	var sayilar [5]int
	var sehirler [5]string

	sayilar[2] = 50
	sayilar[3] = 24
	sayilar[0] = 10
	sehirler[0] = "izmir"
	sehirler[1] = "istanbul"
	sehirler[2] = "ankara"
	sehirler[3] = "konya"
	sehirler[4] = "muğla"

	fmt.Println(sayilar[3])
	fmt.Println(sayilar[2])
	fmt.Println(sayilar)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])

	}
	enbuyuk := 0
	sayilar2 := [5]int{5, 6, 1, 2, 76}
	fmt.Println(sayilar2)
	for i := 0; i < len(sayilar2); i++ {
		if enbuyuk < sayilar2[i] {
			enbuyuk = sayilar2[i]

		}

	}
	fmt.Println(enbuyuk)

	var sayilar3 [2][3]int

	sayilar3[0][0] = 1
	sayilar3[0][1] = 2
	sayilar3[1][0] = 3
	sayilar3[1][1] = 4

	fmt.Println(sayilar3[0][0])
}
