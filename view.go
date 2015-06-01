package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/nsf/termbox-go"

	//"strconv"
)

// Colors from Termbox
const backgroundColor = termbox.ColorBlue
const boardColor = termbox.ColorBlack
const instructionsColor = termbox.ColorYellow

var pieceColors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
}

const defaultMarginWidth = 1
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const boardWidth = 20
const boardHeight = 20
const cellWidth = 2
const boardEndX = boardStartX + boardWidth*cellWidth
const boardEndY = boardStartY + boardHeight
const instructionsStartX = boardEndX + defaultMarginWidth
const instructionsStartY = boardStartY

const title = "Snake Written in Go"

var instructions = []string{

	"Goal: Collect all the dots!",
	"",
	"\u2190		Left",
	"\u2192		Right",
	"\u2191		Up",
	"\u2193		Down",
	"s		Start",
	"esc	Exit",
	"",
	"Dots: %v",
	"",
	"GAME OVER",
}

func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	tbprint(titleStartX, titleStartY, instructionsColor, backgroundColor, title)
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			cellValue := g.board[y][x]
			absCellValue := int(math.Abs(float64(cellValue)))
			cellColor := pieceColors[absCellValue]
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(boardStartX+cellWidth*(x)+i, boardStartY+y, ' ', cellColor, cellColor)
			}
		}
	}
	for y, instruction := range instructions {
		if strings.HasPrefix(instruction, "Dots:") {
			instruction = fmt.Sprintf(instruction, g.dots)
		} else if strings.HasPrefix(instruction, "GAME OVER") && g.state != gameOver {
			instruction = ""
		}
		tbprint(instructionsStartX, instructionsStartY+y, instructionsColor, backgroundColor, instruction)
	}
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
