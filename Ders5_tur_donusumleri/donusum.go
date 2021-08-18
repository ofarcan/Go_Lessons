package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a float32 = 100.7878
	var c string = "-45s"
	// var b int16 = int16(c)   /* strinden integera cevrilecekse harf olmamalı yoksa 0 değeri döner */

	b, _ := strconv.Atoi(c)

	fmt.Printf("a tür: %T , a deger: %v\n", a, a)
	fmt.Printf("b tur: %T , b deger: %v", b, b)

}
