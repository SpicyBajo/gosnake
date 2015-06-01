package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

const animationSpeed = 10 * time.Millisecond

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g := NewGame()
	render(g)

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					g.moveLeft()
				case ev.Key == termbox.KeyArrowRight:
					g.moveRight()
				case ev.Key == termbox.KeyArrowUp:
					g.moveUp()
				case ev.Key == termbox.KeyArrowDown:
					g.moveDown()
				//case ev.Key == termbox.KeySpace:
				//g.moveDown()
				//g.dots++
				case ev.Ch == 's':
					g.start()
					g.play()
				//case ev.Ch == 'p':
				//g.pause()
				case ev.Key == termbox.KeyEsc:
					return
				}
			}
		case <-g.fallingTimer.C:
			g.play()
		default:
			render(g)
			time.Sleep(animationSpeed)
		}
	}

}
