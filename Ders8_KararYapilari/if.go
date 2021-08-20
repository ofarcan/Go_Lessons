package main

import "fmt"

func main() {
	var (
		dizi    [10]int
		tek     []int
		cift    []int
		negatif []int
	)

	fmt.Println("10 tane sayi giriniz: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&dizi[i])

		if dizi[i]%2 == 0 && dizi[i] > 0 {

			cift = append(cift, dizi[i])

		} else if dizi[i]%2 != 0 && dizi[i] > 0 {

			tek = append(tek, dizi[i])

		} else {
			negatif = append(negatif, dizi[i])
		}
	}

	fmt.Println("Girilen sayılar: ", dizi)
	fmt.Println("Tek sayılar: ", tek)
	fmt.Println("Çift Sayılar: ", cift)
	fmt.Println("Negatif sayılar", negatif)

}
