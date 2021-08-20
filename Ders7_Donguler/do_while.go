package main

import "fmt"

func main() {

	var faktoriyel int
	toplam, i := 1, 1

	for {
		fmt.Println("faktoriyel değişkeni giriniz: ")
		fmt.Scan(&faktoriyel)

		if faktoriyel < 0 {
			fmt.Println("Lütfen bir doğal sayı giriniz.")
			break

		}

		for i <= faktoriyel {
			toplam = toplam * i
			i++
		}

		fmt.Println(toplam)
		break

	}

}
