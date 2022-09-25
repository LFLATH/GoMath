package main

import (
	"fmt"

	"github.com/LFLATH/GoMath/matrix"
)

func main() {
	fmt.Println("")
	a := [][]float64{
		{0, 1},
		{4, 5},
	}
	b := [][]float64{
		{0, 1},
		{4, 5},
	}
	c, err := matrix.Dot(a, b)
	if err != nil {
		panic(err)
	}
	matrix.Display(c)
}
