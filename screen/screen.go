/*
Copyright © 2023 tienvu461@gmail.com
*/
package screen

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type screen struct {
	logMsg string
}

var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorBlue,
	termbox.ColorCyan,
	termbox.ColorYellow,
	termbox.ColorWhite,
	termbox.ColorMagenta,
	termbox.ColorLightGray,
	termbox.ColorRed,
}

func (s *screen) Logtb(msg string) {
	s.logMsg = msg
}

func tbprint(msg string) {
	w, h := termbox.Size()
	midy := h / 2
	midx := (w - 30) / 2
	fg := colors[0]
	bg := colors[4]
	for _, c := range msg {
		termbox.SetCell(midx, midy, c, fg, bg)
		midx += runewidth.RuneWidth(c)
	}
}

func (s *screen) RenderAsciiBoard(board [][]int) {
	fmt.Printf("\n%s\n", strings.Repeat("=", len(board[0])))
	for _, e := range board {
		for _, num := range e {
			if num > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func (s *screen) Render(board [][]int) {
	termbox.Clear(termbox.ColorGreen, termbox.ColorGreen)
	offset := 2
	cellWidth := 2
	for y, e := range board {
		for x, num := range e {
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell((x+offset)*cellWidth+i, y+offset, ' ', colors[num], colors[num])
			}
		}
	}
	if s.logMsg != "" {
		tbprint(s.logMsg)
	}
	termbox.Flush()
}

func NewScreen() *screen {
	return &screen{}
}
