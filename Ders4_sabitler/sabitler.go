package main

import (
	"fmt"
)

var b string

// const deger_almayacak string   //constantlar mutlaka başta değer beklerler. yalnızca variable türü ile tanımlama yapılamaz

func main() {

	b = "bu değerdir"
	const a = "bu sabittir" //constant kısaltılması
	fmt.Println(a, "\n", b)

}
