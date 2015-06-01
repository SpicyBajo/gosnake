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
	g.prevLocations = [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	g.dx = []int{1, 0, 0, 0}
	g.dy = []int{1, 0, 0, 0}
	g.dotX = -500
	g.dotY = -500

	g.fallingTimer = time.NewTimer(time.Duration(1000000 * time.Second))
	g.fallingTimer.Stop()

}

//Called every time Play() is called
func (g *Game) resetFallingTimer() {
	g.fallingTimer.Reset(g.speed())
}

//Determines speed based on Dots -- Needs work
func (g *Game) speed() time.Duration {

	multiplier := 0

	switch {
	case g.dots <= 4:
		multiplier = 50
	case g.dots > 4 && g.dots <= 9:
		multiplier = 45
	case g.dots > 9 && g.dots <= 12:
		multiplier = 35
	case g.dots > 12 && g.dots <= 14:
		multiplier = 30
	case g.dots > 14 && g.dots <= 20:
		multiplier = 22
	case g.dots > 20:
		multiplier = 25
	}

	return 500*time.Millisecond - (time.Duration(g.dots*multiplier) * time.Millisecond)

}

//This goes off for every timer tick
func (g *Game) play() {

	//Ends game if snake is outside boarders
	if g.x < 0 || g.x > 19 || g.y < 0 || g.y > 19 {

		g.state = gameOver
		g.resetGame()
		g.start()

	}

	g.erasePiece()

	if g.d > 0 {

		// Moves snake body around by recording previous locations of snake head
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

	//Moves snake head beased on current direction
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

	//Ends game if snake hits itself
	for i := g.d; i > 0; i-- {

		if g.x == g.prevLocations[i][1] && g.y == g.prevLocations[i][0] {

			g.state = gameOver
			g.resetGame()
			g.start()

		}

	}

	//If snake head its dot, draws new dot and continues game play
	if g.x == (g.dotLocation[0]) && g.y == (g.dotLocation[1]) {

		g.resetFallingTimer()
		g.dots++
		g.d++
		g.getDot()
		g.play()

	}

	g.fillMatrix()

}

//Gets a single snake head
func (g *Game) getPiece() bool {
	g.piece = 1
	g.x = boardWidth / 2
	g.y = boardHeight / 2

	g.dx[0] = 0
	g.dy[0] = 0

	return true
}

//Finds new position for dot when it is eaten by snake
func (g *Game) getDot() bool {

	g.dot = 1

	g.dotX = rand.Int() % boardWidth
	g.dotY = rand.Int() % boardHeight

	for i := 0; i < g.dots; i++ {

		if g.dotX == g.prevLocations[i][1] && g.dotY == g.prevLocations[i][0] {

			g.getDot()

		}

	}

	g.dotLocation[0] = g.dotX
	g.dotLocation[1] = g.dotY

	return true
}

//Draws everything on board
func (g *Game) fillMatrix() {
	// Fill Snake Head
	for k := 0; k < numSquares; k++ {
		x := g.x
		y := g.y

		g.board[y][x] = 1

	}
	// fill dot
	for k := 0; k < numSquares; k++ {
		x := g.dotX
		y := g.dotY

		g.board[y][x] = 2

	}
	// fill extra Snake
	for k := 0; k < numSquares; k++ {

		for m := 0; m < g.d; m++ {

			x := g.prevLocations[m][1]
			y := g.prevLocations[m][0]

			g.board[y][x] = 1

		}

	}

	g.resetFallingTimer()

}

//Draws everything in new positions for every time tick
func (g *Game) placePiece() {

	for k := 0; k < numSquares; k++ {
		x := g.x + g.dx[k]
		y := g.y + g.dy[k]
		if 0 <= y && y < boardHeight && 0 <= x && x < boardWidth {
			g.board[y][x] = -g.piece
		}
	}

}

//Clears board for re-draw every time tick
func (g *Game) erasePiece() {
	//erase snake head
	for k := 0; k < numSquares; k++ {
		x := g.x + g.dx[k]
		y := g.y + g.dy[k]
		if 0 <= y && y < boardHeight && 0 <= x && x < boardWidth {
			g.board[y][x] = 0
		}

	}
	//erase extra snake
	for k := 0; k < numSquares; k++ {

		for m := 0; m < (g.d); m++ {
			x := g.prevLocations[m][1]
			y := g.prevLocations[m][0]

			g.board[y][x] = 0

		}
	}
}

//User pressed S to start game
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
