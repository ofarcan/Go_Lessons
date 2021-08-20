package main

import "fmt"

func main() {

	var faktoriyel int
	toplam, i := 1, 1

	fmt.Println("faktoriyel değişkeni giriniz: ")
	fmt.Scan(&faktoriyel) //scan fonksiyonu kullanırken & unutulmamalıdır.

	for i <= faktoriyel { // aslında go dilinde while dongusu yok for ile bu şekilde oluşturulur.
		toplam = toplam * i
		i++
	}

	fmt.Println(toplam)

}
