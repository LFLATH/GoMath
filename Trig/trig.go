package trig

import (
	"math"
)

func CosH(x float64) float64 {
	e := 2.718281828459045
	return ((math.Pow(e, x) + math.Pow(e, -x)) / 2)
}

func SinH(x float64) float64 {
	e := 2.718281828459045
	return ((math.Pow(e, x) - math.Pow(e, -x)) / 2)
}

func TanH(x float64) float64 {
	return (SinH(x) / CosH(x))

}
