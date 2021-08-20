package main

import "fmt"

func main() {

	/*
		toplam := 0
		for i := 0; i < 10; i++ {
			toplam = toplam + i										//toplama örneği
		}

		fmt.Println(toplam)
	*/

	///////////////////////////////////////

	faktoriyel := 0
	fmt.Println("Faktöriyel değişkeni giriniz: ")
	fmt.Scan(&faktoriyel)

	toplam := 1
	for i := 1; i <= faktoriyel; i++ { //faktoriyel hesaplama
		toplam = toplam * i
	}

	fmt.Println(toplam)

	///////////////////////////while dongusu //////////////////////////////
}
