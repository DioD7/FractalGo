package program

import (
	"fmt"
	"time"
)

var (
	xMin = -1.5
	xMax = 0.5
	yMin = -1.0
	yMax = 1.0
)

type Pixel struct {
	x    int
	y    int
	iter int
}

func setBounds(x, y int) {
	xLen := float64(xMax-xMin) / (2.0 * zoom)
	yLen := float64(yMax-yMin) / (2.0 * zoom)
	xStep := float64(xMax-xMin) / screenWidth
	yStep := float64(yMax-yMin) / screenHeight
	xTrue := xMin + float64(x)*xStep
	yTrue := yMin + float64(y)*yStep

	xMin = xTrue - xLen
	xMax = xTrue + xLen
	yMin = yTrue - yLen
	yMax = yTrue + yLen
}

func generateImage() {
	fmt.Println("Generating image")
	start := time.Now()
	xStep := float64(xMax-xMin) / screenWidth
	yStep := float64(yMax-yMin) / screenHeight
	for i := 0; i < screenHeight; i++ {
		for j := 0; j < screenWidth; j++ {
			value := generatePixel(xMin+xStep*float64(j), yMin+yStep*float64(i))
			r, g, b := color(value)
			currentPixel := (i*screenWidth + j) * 4
			imgPix[currentPixel] = r
			imgPix[currentPixel+1] = g
			imgPix[currentPixel+2] = b
			imgPix[currentPixel+3] = 0xff
		}
	}
	duration := time.Since(start)
	fmt.Println("[Finished in", duration, "]")
	img.ReplacePixels(imgPix)
}

func color(it int) (r, g, b byte) {
	if it == maxIter {
		return 0xff, 0xff, 0xff
	}
	c := palette[it]
	return c, c, c
}

func generatePixel(x, y float64) int {
	z := complex(0, 0)
	c := complex(x, y)
	it := 0
	for ; it < maxIter; it++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
	}
	return it
}
