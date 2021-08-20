package main

import "fmt"

func main() {

	// Go dilinde for each yerine For Range dongusu mevcut.

	slice := []string{"omer", "faruk", "can"}
	/*
		for a := range slice {	//yalnızca indexlerini döndürür
			fmt.Println(a)
		}
	*/

	for i, a := range slice { //hem index hem string değerleri döndürür
		fmt.Println(i, a)
	}

	/*
		for _, a := range slice { // indexleri döndürmeden yalnızca string değerleri döndürdü.
			fmt.Println(a)
		}
	*/

	/*
		for i := 0; i < len(slice); i++ {		//len dizi uzunluğunu aldı.

			fmt.Println(slice[i])			//for range yerine yalnızca for dongusu ile de yazılabilir.
		}

	*/

}
