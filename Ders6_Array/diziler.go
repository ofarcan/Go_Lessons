package main

import "fmt"

func main() {

	/* diziler //slice // map */

	//Diziler
	var dizi [2]string
	dizi[0] = "ömer"
	dizi[1] = "faruk"

	fmt.Println(dizi)

	//Slice
	slice := make([]int, 2)
	slice[1] = 1
	slice = append(slice, 2, 3, 4, 5)
	fmt.Println(slice)

	var slice2 []int
	slice2 = slice[2:5]
	fmt.Println(slice2)

	//map

	map1 := make(map[string]string)
	map1["ad"] = "cristiano"
	map1["soyad"] = "ronaldo"

	fmt.Println(map1)

}
