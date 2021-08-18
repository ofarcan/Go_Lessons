package main

import "fmt"

func main() {

	var a int = 6
	b := "merhaba"
	var c, d = "dunya", false
	d, e := true, 5.2

	fmt.Println(a, b, c, d, e)
	fmt.Printf("Tür: %T Değer: %v\n", e, a)

}
