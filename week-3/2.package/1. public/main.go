package main

import (
	"fmt"

	"github.com/nsgnd/dersnotlar-picus-go/week-3/2. package/1. public/formatter"
	"github.com/nsgnd/dersnotlar-picus-go/week-3/2. package/1. public/maths"
)

func main() {
	num := maths.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
