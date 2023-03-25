/*
Copyright © 2023 tienvu461@gmail.com
*/
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/tienvu461/gotris/screen"
	"github.com/tienvu461/gotris/tetris"
)

const FPS = 60 * time.Millisecond

func main() {
	rand.Seed(time.Now().UnixNano())
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

	ticker := time.NewTimer(time.Duration(FPS))
	game := tetris.NewGame()
	screen := screen.NewScreen()
	for game.GetState() != tetris.G_OVER {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowUp:
					game.Rotate()
				case ev.Key == termbox.KeyArrowDown:
					game.SpeedUp()
					// game.MoveDown()
				case ev.Key == termbox.KeyArrowLeft:
					game.MoveLeft()
				case ev.Key == termbox.KeyArrowRight:
					game.MoveRight()
				case ev.Key == termbox.KeySpace:
					game.Fall()
				case ev.Ch == 'q':
					// quit
					return
				case ev.Ch == 's':
					//start
					game.Start()
				default:
					return
				}
			}
		case <-ticker.C:
			screen.Render(game.GetBoard())
			screen.Logtb(fmt.Sprintf("Score: %d\n", game.GetScore()))
			ticker.Reset(FPS)
		case <-game.FallSpeed.C:
			game.GameLoop()
		}
	}
}
