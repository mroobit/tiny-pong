package main

import (
	"time"

	"github.com/tinygo-org/tinygo/src/machine"
	"tinygo.org/x/drivers/microbitmatrix"
)

func main() {
	// set LED display
	display := microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	// set buttons
	bta := machine.BUTTONA
	bta.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	btb := machine.BUTTONB
	btb.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	// set players, ball
	h := NewPlayer(0, 2)
	h.ball = true
	m := NewPlayer(4, 2)
	b := NewBall()

	// display opening splash

	// display menu

	// game loop
	t := time.Now()
	for h.score < 5 && m.score < 5 {
		if time.Since(t).Microseconds() >= 100000 {
			// reset time back
			t = time.Now()

			// check if ball is offscreen
			if b.Offscreen() {
				// reset player positions
				m.loc[1] = 2
				h.loc[1] = 2
				// add to score, assign ball to player who scored
				if b.loc[0] == -1 {
					m.score++
					m.Get(b)
					b.loc[0] = 3
				}
				if b.loc[0] == 5 {
					h.score++
					h.Get(b)
					b.loc[0] = 1
				}
			}

			// if ball isn't held, move the ball
			if !m.ball && !h.ball {
				b.Move()
				// if ball hits paddle, reverse both dirs
				b.Hit(m, h)
			}

			// move machine player, release ball if machine carrying
			m.Move()
			if m.ball == true {
				m.Carry(b)
				m.ball = false
			}

			// if button b is pressed, release ball
			if !btb.Get() {
				h.ball = false
			}

			// if button a is pressed, move human player
			if !bta.Get() {
				h.Move()
				if h.ball == true {
					h.Carry(b)
				}
			}

			// clear display, set pixelsfor both paddles and ball
			display.ClearDisplay()
			display.SetPixel(h.loc[0], h.loc[1], microbitmatrix.BrightnessFull)
			display.SetPixel(m.loc[0], m.loc[1], microbitmatrix.BrightnessFull)
			display.SetPixel(b.loc[0], b.loc[1], microbitmatrix.BrightnessFull)
		}
		// display pixel grid
		display.Display()
	}
	// after loop
	// set score pixels
	// display pixel grid

	for {
		display.ClearDisplay()
		for i := 0; i < int(h.score); i++ {
			display.SetPixel(1, int16(i), microbitmatrix.BrightnessFull)
		}
		for i := 0; i < int(m.score); i++ {
			display.SetPixel(3, int16(i), microbitmatrix.BrightnessFull)
		}
		display.Display()
	}

}

type Player struct {
	loc   [2]int16
	dir   int16
	score int16
	ball  bool
}

func NewPlayer(x, y int16) *Player {
	p := &Player{
		loc: [2]int16{x, y},
		dir: 1,
	}
	return p
}

func (p *Player) Move() {
	if p.loc[1] == 0 || p.loc[1] == 4 {
		p.dir *= -1
	}
	p.loc[1] += p.dir
}

func (p *Player) Carry(b *Ball) {
	if p.loc[1] == 0 || p.loc[1] == 4 {
		b.dir[1] = p.dir * -1
	}
	b.loc[1] += b.dir[1]
}

func (p *Player) Get(b *Ball) {
	p.ball = true
	b.loc[1] = p.loc[1] + p.dir
	b.dir[1] = p.dir
}

func (p *Player) Reset() {
	p.loc[1] = 2
}

func (p *Player) Score() {
	p.score++
}

type Ball struct {
	loc [2]int16
	dir [2]int16
}

func NewBall() *Ball {
	b := &Ball{
		loc: [2]int16{1, 3},
		dir: [2]int16{1, 1},
	}
	return b
}

func (b *Ball) Move() {
	if b.loc[1] == 0 || b.loc[1] == 4 {
		b.dir[1] *= -1
	}
	b.loc[0] += b.dir[0]
	b.loc[1] += b.dir[1]
}

func (b *Ball) Hit(m, h *Player) {
	if (b.loc[0] == 1 && b.loc[1] == h.loc[1]) ||
		(b.loc[0] == 3 && b.loc[1] == m.loc[1]) {
		b.dir[0] *= -1
		b.dir[1] *= -1
	}
}

func (b *Ball) Offscreen() bool {
	if b.loc[0] == -1 || b.loc[0] == 5 {
		return true
	}
	return false
}
