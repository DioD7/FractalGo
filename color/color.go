package color

import (
	"errors"
	"fmt"
	"math"
)

type Color struct {
	R int
	G int
	B int
}

var (
	BLACK   = Color{0, 0, 0}
	WHITE   = Color{255, 255, 255}
	RED     = Color{255, 0, 0}
	GREEN   = Color{0, 255, 0}
	BLUE    = Color{0, 0, 255}
	YEALLOW = Color{255, 255, 0}
	CYAN    = Color{0, 255, 255}
	PURPLE  = Color{255, 0, 255}

	Palette []Color

	factVal uint64 = 1
)

func SetPalette(p []Color) {
	Palette = p
}

func GetColor(t float64) (Color, error) {
	if len(Palette) == 0 {
		return BLACK, errors.New("palette not specified")
	}
	var r float64 = 0.0
	var g float64 = 0.0
	var b float64 = 0.0
	n := len(Palette) - 1
	for i := 0; i < len(Palette); i++ {
		multiplicator := float64(combination(n, i)) * math.Pow(t, float64(i)) * math.Pow(1.0-t, float64(n-i))
		r += multiplicator * float64(Palette[i].R)
		g += multiplicator * float64(Palette[i].G)
		b += multiplicator * float64(Palette[i].B)
	}
	rezult := Color{int(r), int(g), int(b)}
	return rezult, nil
}

func combination(n, i int) uint64 {
	answer := factorial(n) / (factorial(i) * factorial(n-i))
	return answer
}

//Factorial function taken from https://www.golangprograms.com/go-program-to-find-factorial-of-a-number.html
func factorial(n int) uint64 {
	factVal = 1
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}
