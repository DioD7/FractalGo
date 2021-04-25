package program

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const (
	screenWidth  = 700
	screenHeight = 700
	maxIter      = 500
)

var (
	img     = ebiten.NewImage(screenWidth, screenHeight)
	imgPix  []byte
	palette [maxIter]byte
	xMouse  int
	yMouse  int

	xMin = -1.5
	xMax = 0.5
	yMin = -1.0
	yMax = 1.0

	mousePressed = false
)

func setBounds(x, y int) {
	xLen := float64(xMax-xMin) / 4.0
	yLen := float64(yMax-yMin) / 4.0
	xStep := float64(xMax-xMin) / screenWidth
	yStep := float64(yMax-yMin) / screenHeight
	xTrue := xMin + float64(x)*xStep
	yTrue := yMin + float64(y)*yStep

	xMin = xTrue - xLen
	xMax = xTrue + xLen
	yMin = yTrue - yLen
	yMax = yTrue + yLen
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mousePressed = true
		xMouse, yMouse = ebiten.CursorPosition()
	} else if mousePressed {
		mousePressed = false
		setBounds(xMouse, yMouse)
		generateImage()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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

func color(it int) (r, g, b byte) {
	if it == maxIter {
		return 0xff, 0xff, 0xff
	}
	c := palette[it]
	return c, c, c
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

func initialize() {
	fmt.Println("Initilaizing")
	imgPix = make([]byte, 4*screenWidth*screenHeight)
	for i := range palette {
		palette[i] = byte(math.Sqrt(float64(i)/float64(len(palette))) * maxIter)
	}
	generateImage()
}

func Launch() {
	fmt.Println("Launching fractal visualization")
	initialize()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Fractal viewer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
