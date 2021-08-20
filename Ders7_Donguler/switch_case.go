package main

import "fmt"

func main() {

	var gun int
	fmt.Println("kacıncı gundesiniz")
	fmt.Scan(&gun)

	switch gun {
	case 1:
		fmt.Printf("Pazartesi")
	case 2:
		fmt.Printf("Salı")
	case 3:
		fmt.Printf("Çarşamba")
	case 4:
		fmt.Printf("Perşembe")
	case 5:
		fmt.Printf("Cuma")
	case 6:
		fmt.Printf("Ctesi")
	case 7:
		fmt.Printf("Pazar")

	default:
		fmt.Print("Tanımlanamayan bir deger girdiniz.")

	}

}
