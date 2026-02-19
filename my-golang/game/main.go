package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 400
	screenHeight = 400
	cellSize     = 20
	gridSize     = screenWidth / cellSize
)

type Point struct {
	x, y int
}

type Game struct {
	snake     []Point
	dir       Point
	food      Point
	gameOver  bool
	tickCount int
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	return &Game{
		snake: []Point{{5, 5}},
		dir:   Point{1, 0},
		food:  Point{rand.Intn(gridSize), rand.Intn(gridSize)},
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	// Controls
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.dir.y != 1 {
		g.dir = Point{0, -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.dir.y != -1 {
		g.dir = Point{0, 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.dir.x != 1 {
		g.dir = Point{-1, 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.dir.x != -1 {
		g.dir = Point{1, 0}
	}

	// Slow movement
	g.tickCount++
	if g.tickCount < 10 {
		return nil
	}
	g.tickCount = 0

	head := g.snake[0]
	newHead := Point{head.x + g.dir.x, head.y + g.dir.y}

	// Wall collision
	if newHead.x < 0 || newHead.y < 0 || newHead.x >= gridSize || newHead.y >= gridSize {
		g.gameOver = true
		return nil
	}

	// Self collision
	for _, s := range g.snake {
		if s == newHead {
			g.gameOver = true
			return nil
		}
	}

	// Move snake
	g.snake = append([]Point{newHead}, g.snake...)

	// Eat food
	if newHead == g.food {
		g.food = Point{rand.Intn(gridSize), rand.Intn(gridSize)}
	} else {
		g.snake = g.snake[:len(g.snake)-1]
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Draw snake
	for _, s := range g.snake {
		ebitenutil.DrawRect(screen,
			float64(s.x*cellSize),
			float64(s.y*cellSize),
			cellSize, cellSize,
			color.RGBA{0, 255, 0, 255})
	}

	// Draw food
	ebitenutil.DrawRect(screen,
		float64(g.food.x*cellSize),
		float64(g.food.y*cellSize),
		cellSize, cellSize,
		color.RGBA{255, 0, 0, 255})

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER")
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game (Go)")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
