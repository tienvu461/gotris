/*
Copyright © 2023 tienvu461@gmail.com
*/
package screen

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
)

type screen struct {
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
	termbox.Flush()
}

func NewScreen() *screen {
	return &screen{}
}
