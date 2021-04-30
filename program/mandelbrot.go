package program

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/DioD7/fractal/color"
)

const (
	numProcesses = 10
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
	//Parallel
	var wg sync.WaitGroup
	out := make(chan Pixel)
	deltax := screenWidth / numProcesses
	for p := 0; p < numProcesses; p++ {
		procStart := p * deltax
		wg.Add(1)
		go func(pStart int) {
			defer wg.Done()
			generateSubImage(pStart, deltax, out)
		}(procStart)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	for px := range out {
		currentPixel := (px.y*screenWidth + px.x) * 4
		r, g, b := getColor(px.iter)
		imgPix[currentPixel] = r
		imgPix[currentPixel+1] = g
		imgPix[currentPixel+2] = b
		imgPix[currentPixel+3] = 0xff
	}
	duration := time.Since(start)
	fmt.Println("[Finished in", duration, "]")
	img.ReplacePixels(imgPix)
}

func generateSubImage(start, len int, out chan Pixel) {
	xStep := float64(xMax-xMin) / screenWidth
	yStep := float64(yMax-yMin) / screenHeight
	for i := 0; i < screenHeight; i++ {
		for j := start; j < start+len; j++ {
			value := generatePixel(xMin+xStep*float64(j), yMin+yStep*float64(i))
			out <- Pixel{j, i, value}
		}
	}
}

func getColor(it int) (r, g, b byte) {
	t := math.Sqrt(float64(it) / float64(maxIter))
	c, _ := color.GetColor(t)
	return byte(c.R), byte(c.G), byte(c.B)
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
