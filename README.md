# gotris

Tetris written in go
Can be play on terminal

## How to

To run using golang:
`go run main.go`

To run using binary:
`./gotris`

To build binary:
`make -f build.mk build`

## Controls:

| #   | Key   | Action     |
| --- | ----- | ---------- |
| 1   | S     | Start      |
| 2   | Q     | Quit       |
| 3   | ←     | Move left  |
| 4   | →     | Move right |
| 5   | ↑     | Rotate     |
| 6   | ↓     | Soft drop  |
| 7   | Space | Hard drop  |

## Sceenshots

![Screenshot](screen_example.png?raw=true 'Screenshot')

## Logic

Rotate rule

![Rotate](rotate-pieces.webp?raw=true 'rotate_rule')

## TODOs:

- [ ] Welcome, gameover, leaderboard, screen
- [ ] Configurable size, speed
- [ ] Using WebAssembly to be playable on browsers
