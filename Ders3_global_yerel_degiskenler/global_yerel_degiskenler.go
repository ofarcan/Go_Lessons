package main

import "fmt"

var b int = 11 //global değişken, fonksiyon dışında tanımlandığı için

func main() {

	a := 15 //yerel değişkenler

	fmt.Println(a, b)
	bar()

}

func bar() {
	c := 19
	fmt.Println(c)
}
