package main

import (
	"fmt"
	"math"
)

var (
	GlobalAlan = 100 // global ise büyük harf ile belirtilir
)

var Global2 = "Global"

func main() {
	fmt.Println(GlobalAlan)

	sum := GlobalAlan + 10
	fmt.Printf("%v %T\n", sum, sum)

	toplaMutlak := math.Abs(float64(GlobalAlan))
	fmt.Printf("%v %T\n", toplaMutlak, toplaMutlak)

	fmt.Println(Global2)
}
