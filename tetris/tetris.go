/*
Copyright © 2023 tienvu461@gmail.com
*/
package tetris

import "time"

const B_HEIGHT = 10
const B_WIDTH = 9

type gameState int

const (
	gameInit gameState = iota
	gamePlay
	gameOver
)

type game struct {
	board     [][]int
	position  vector
	block     block
	state     gameState
	FallSpeed *time.Timer
}

func (g *game) genBlock() {
	g.block = randBlock()
	g.position = vector{1, B_WIDTH / 2}
}

func (g *game) resetFallSpeed() {
	g.FallSpeed.Reset(700 * time.Millisecond)
}

func (g *game) Start() {
	g.state = gamePlay
	g.genBlock()
	g.resetFallSpeed()
}

func (g *game) colision() bool {
	for _, v := range g.block.shape {
		pos := g.blockOnBoardByPosition(v)
		if pos.x < 0 || pos.x >= B_WIDTH {
			return true
		}
		if pos.y >= B_HEIGHT {
			return true
		}
		if g.board[pos.y][pos.x] > 0 {
			return true
		}
	}
	return false
}

func (g *game) moveIfPosible(v vector) bool {
	g.position.x += v.x
	g.position.y += v.y
	if g.colision() {
		g.position.x -= v.x
		g.position.y -= v.y
		return false
	}
	return true
}

func (g *game) MoveLeft() {
	g.moveIfPosible(vector{0, -1})
}

func (g *game) MoveRight() {
	g.moveIfPosible(vector{0, 1})
}

func (g *game) SpeedUp() {
	g.FallSpeed.Reset(50 * time.Millisecond)
}

func (g *game) Rotate() {
	g.block.Rotate()
	// TODO: handle exception rotate will crash on border
}

func (g *game) Fall() {
	for {
		if !g.moveIfPosible(vector{1, 0}) {
			g.FallSpeed.Reset(1 * time.Millisecond)
			return
		}
	}
}
func (g *game) lockBlocks() {
	g.board = g.GetBoard()
}
func (g *game) clearLine() {
	line := make([]int, B_WIDTH)
	for i := 0; i < B_WIDTH; i++ {
		line[i] = 0
	}
	clearLine := [][]int{line}
	for y := 0; y < B_HEIGHT; y++ {
		for x := 0; x < B_WIDTH; x++ {
			if g.board[y][x] == 0 {
				break
			} else if x == B_WIDTH-1 {
				newBoard := append(clearLine, g.board[:y]...)
				g.board = append(newBoard, g.board[y+1:]...)
			}
		}
	}
}
func (g *game) GameLoop() {
	if !g.moveIfPosible(vector{1, 0}) {
		g.lockBlocks()
		g.clearLine()
		g.genBlock()
		if g.colision() {
			g.FallSpeed.Stop()
			g.state = gameOver
			return
		}
	}
	g.resetFallSpeed()
}

func (g *game) blockOnBoardByPosition(v vector) vector {
	px := g.position.x + v.x
	py := g.position.y + v.y
	return vector{py, px}
}

// return current board state
func (g *game) GetBoard() [][]int {
	cBoard := make([][]int, len(g.board))
	for y := 0; y < len(g.board); y++ {
		cBoard[y] = make([]int, len(g.board[y]))
		copy(cBoard[y], g.board[y])
		// for x := 0; x < len(g.board[y]); x++ {
		// 	cBoard[y][x] = g.board[y][x]
		// }
	}

	for _, v := range g.block.shape {
		pos := g.blockOnBoardByPosition(v)
		cBoard[pos.y][pos.x] = g.block.color
	}
	return cBoard
}

func (g *game) init() {
	// initialize 2d array
	g.board = make([][]int, B_HEIGHT)
	for y := 0; y < B_HEIGHT; y++ {
		g.board[y] = make([]int, B_WIDTH)
		for x := 0; x < B_WIDTH; x++ {
			g.board[y][x] = 0
		}
	}
	g.position = vector{0, B_WIDTH / 2}
	// g.block = blocks[0]
	g.FallSpeed = time.NewTimer(time.Duration(1000 * time.Second))
	g.FallSpeed.Stop()
	g.state = gameInit
}

func NewGame() *game {
	g := &game{}
	g.init()
	return g
}
