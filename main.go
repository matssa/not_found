package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Running http server
// func main() {
// 	http.Handle("/", http.FileServer(http.Dir(".")))
// 	http.ListenAndServe(":3000", nil)
// }

type Direction int32

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

type Snake struct {
	direction Direction
	Head      [2]int
	Body      int
}

type Game struct {
	snake Snake
	score int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.snake.direction = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.snake.direction = Right
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.snake.direction = Down
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.snake.direction = Left
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.snake.direction == Up {
		ebitenutil.DebugPrint(screen, "Snake direction: Up")
	} else if g.snake.direction == Right {
		ebitenutil.DebugPrint(screen, "Snake direction: Right")
	} else if g.snake.direction == Down {
		ebitenutil.DebugPrint(screen, "Snake direction: Down")
	} else {
		ebitenutil.DebugPrint(screen, "Snake direction: Left")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func SetGameMeta(x int, y int) error {
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("Snake")
	return nil
}

func (g *Game) InitGame(x int, y int) error {
	g.snake.direction = Up
	g.snake.Head = [2]int{x, y}
	return nil
}

func main() {
	g := &Game{
		score: 0,
	}

	g.InitGame(500, 500)

	SetGameMeta(500, 500)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
