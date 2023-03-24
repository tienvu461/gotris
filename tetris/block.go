/*
Copyright © 2023 tienvu461@gmail.com
*/
package tetris

import (
	"math"
	"math/rand"
)

type vector struct {
	y, x int
}
type block struct {
	shape []vector
	color int
}

func (b *block) Rotate() {
	var rotateAngle (float64) = 90
	rotateRadian := rotateAngle * math.Pi / 180
	cos := int(math.Round(math.Cos(rotateRadian)))
	sin := int(math.Round(math.Sin(rotateRadian)))

	for i, e := range b.shape {
		ny := e.y*cos - e.x*sin
		nx := e.y*sin - e.x*cos

		b.shape[i] = vector{ny, nx}
	}
}

// https://qph.cf2.quoracdn.net/main-qimg-356e2b21c801381db2890dab49a9ea88
var blocks = []block{
	{
		shape: []vector{{0, 0}},
		color: 0,
	},
	// 1. L block - Orange Ricky
	{
		shape: []vector{{0, 1}, {1, 1}, {1, 0}, {1, -1}},
		color: 1,
	},
	// 2. Oposite L block - Blue Ricky
	{
		shape: []vector{{0, -1}, {1, -1}, {1, 0}, {1, 1}},
		color: 2,
	},
	// 3. Z Block - Cleverland Z
	{
		shape: []vector{{0, -1}, {0, 0}, {1, 0}, {1, 1}},
		color: 3,
	},
	// 4. S block - Rhode Island Z
	{
		shape: []vector{{1, -1}, {1, 0}, {0, 0}, {0, 1}},
		color: 4,
	},
	// 5. I Block - Hero
	{
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
		color: 5,
	},
	// 6. Upsidedown T Block - Teewee
	{
		shape: []vector{{1, -1}, {0, 0}, {1, 0}, {1, 1}},
		color: 6,
	},
	// 7. Square Block - Smashboy
	{
		shape: []vector{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		color: 7,
	},
}

func randBlock() block {
	idx := rand.Intn(len(blocks)-1) + 1
	return blocks[idx]
}
