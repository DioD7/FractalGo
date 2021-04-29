package program

import (
	"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const (
	screenWidth  = 700
	screenHeight = 700
	maxIter      = 1000
)

var (
	img     = ebiten.NewImage(screenWidth, screenHeight)
	imgPix  []byte
	palette [maxIter]byte
	xMouse  int
	yMouse  int

	zoom = 2.0

	mousePressed = false
)

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
