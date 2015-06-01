package main

import (
	"math/rand"
	"time"
)

type gameState int

const (
	gameIntro gameState = iota
	gameStarted
	gamePaused
	gameOver
	numSquares = 1
	numTypes   = 7
	slowSpeed  = 700 * time.Millisecond
	HighSpeed  = 150 * time.Millisecond
)

// Pieces
var dxBank = [][]int{
	{},
	{0, 1, -1, 0},
	{0, 1, -1, -1},
	{0, 1, -1, 1},
	{0, -1, 1, 0},
	{0, 1, -1, 0},
	{0, 1, -1, -2},
	{0, 1, 1, 0},
}

var dyBank = [][]int{
	{},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
	{0, 0, 1, 1},
	{0, 0, 1, 1},
	{0, 0, 0, 0},
	{0, 0, 1, 1},
}

type Game struct {
	board         [][]int // [y][x]
	prevLocations [][]int // [x][y]
	state         gameState
	dots          int
	dotLocation   []int // [0] = y [1]=x
	piece         int
	dot           int
	x             int
	y             int
	d             int
	dotX          int
	dotY          int
	direction     int
	dx            []int
	dy            []int
	dxPrime       []int
	dyPrime       []int
	fallingTimer  *time.Timer
}

func NewGame() *Game {
	g := new(Game)
	g.resetGame()
	return g
}

// Reset the game in order to play again.
func (g *Game) resetGame() {
	g.board = make([][]int, boardHeight)
	for y := 0; y < boardHeight; y++ {
		g.board[y] = make([]int, boardWidth)
		for x := 0; x < boardWidth; x++ {
			g.board[y][x] = 0
		}
	}

	g.state = gameIntro
	g.dots = 1
	g.d = 0
	g.x = 5
	g.y = 5
	g.direction = 3
	g.dotLocation = []int{0, 0}
	g.prevLocations = [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	g.dx = []int{1, 0, 0, 0}
	g.dy = []int{1, 0, 0, 0}
	g.dxPrime = []int{1, 0, 0, 0}
	g.dyPrime = []int{1, 0, 0, 0}
	g.dotX = -500
	g.dotY = -500

	g.fallingTimer = time.NewTimer(time.Duration(1000000 * time.Second))
	g.fallingTimer.Stop()

}

func (g *Game) resetFallingTimer() {
	g.fallingTimer.Reset(g.speed())
}

func (g *Game) speed() time.Duration {
	return slowSpeed - HighSpeed*time.Duration(g.dots/2)

}

// Set the timer to make the pieces fall again.

func (g *Game) play() {

	g.erasePiece()

	if g.d > 0 {

		//m := g.d - 1
		//n := m - 1
		//p := m - 1

		/*
			g.prevLocations[q][0] = g.prevLocations[n][0]
			g.prevLocations[q][1] = g.prevLocations[n][1]

			g.prevLocations[n][0] = g.prevLocations[m][0]
			g.prevLocations[n][1] = g.prevLocations[m][1]
		*/

		if g.d > 1 {
			for i := g.d - 1; i > 0; i-- {

				p := i - 1

				g.prevLocations[i][0] = g.prevLocations[p][0]
				g.prevLocations[i][1] = g.prevLocations[p][1]

			}
		}

		g.prevLocations[0][0] = g.y
		g.prevLocations[0][1] = g.x
	}

	switch {
	case g.direction == 1:
		g.x++
	case g.direction == 2:
		g.x--
	case g.direction == 3:
		g.y++
	case g.direction == 4:
		g.y--
	}

	if g.x == (g.dotLocation[0]) && g.y == (g.dotLocation[1]) {
		g.dots++
		g.d++
		g.getDot()
		g.resetFallingTimer()

	}

	g.fillMatrix()
	g.resetFallingTimer()

}

func (g *Game) getPiece() bool {
	g.piece = 1
	g.x = boardWidth / 2
	g.y = boardHeight / 2

	for k := 0; k < numSquares; k++ {
		g.dx[k] = dxBank[g.piece][k]
		g.dy[k] = dyBank[g.piece][k]
	}
	for k := 0; k < numSquares; k++ {
		g.dxPrime[k] = g.dx[k]
		g.dyPrime[k] = g.dy[k]
	}

	return true
}

func (g *Game) getDot() bool {

	g.dot = 1

	g.dotX = rand.Int() % boardWidth
	g.dotY = rand.Int() % boardHeight

	g.dotLocation[0] = g.dotX
	g.dotLocation[1] = g.dotY

	return true
}

func (g *Game) fillMatrix() {
	for k := 0; k < numSquares; k++ { // Fill Snake
		x := g.x
		y := g.y

		g.board[y][x] = 1

	}

	for k := 0; k < numSquares; k++ { // fill DOT
		x := g.dotX
		y := g.dotY

		g.board[y][x] = 2

	}

	for k := 0; k < numSquares; k++ { // fill extra Snake

		for m := 0; m < g.d; m++ {

			x := g.prevLocations[m][1]
			y := g.prevLocations[m][0]

			g.board[y][x] = 1

		}

	}

}

func (g *Game) placePiece() {

	for k := 0; k < numSquares; k++ {
		x := g.x + g.dx[k]
		y := g.y + g.dy[k]
		if 0 <= y && y < boardHeight && 0 <= x && x < boardWidth {
			g.board[y][x] = -g.piece
		}
	}

}

func (g *Game) erasePiece() {
	for k := 0; k < numSquares; k++ {
		x := g.x + g.dx[k]
		y := g.y + g.dy[k]
		if 0 <= y && y < boardHeight && 0 <= x && x < boardWidth {
			g.board[y][x] = 0
		}

	}

	for k := 0; k < numSquares; k++ { // erase extra Snake

		for m := 0; m < (g.d); m++ {
			x := g.prevLocations[m][1]
			y := g.prevLocations[m][0]

			g.board[y][x] = 0

		}
	}
}

func (g *Game) start() {
	switch g.state {
	case gameStarted:
		return
	default:
		g.state = gameStarted
		g.getDot()
		g.getPiece()
		g.placePiece()
		g.fillMatrix()

	}
}

// The user pressed the left arrow.
func (g *Game) moveLeft() {
	if g.state != gameStarted {
		return
	}

	g.direction = 2
	g.play()
}

// The user pressed the right arrow.
func (g *Game) moveRight() {
	if g.state != gameStarted {
		return
	}

	g.direction = 1
	g.play()
}

func (g *Game) moveDown() bool {
	if g.state != gameStarted {
		return false
	}
	g.direction = 3
	g.play()
	return true
}

func (g *Game) moveUp() bool {
	if g.state != gameStarted {
		return false
	}

	g.direction = 4
	g.play()
	return true
}
