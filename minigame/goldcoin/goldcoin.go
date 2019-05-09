package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	boardWidth  = 50
	boardHeight = 15
	goldCount   = 10
)

var (
	point int
	px    int
	golds []*gold
)

type gold struct {
	y     int
	score int
	c     termbox.Attribute
	cd    int
}

func goldRainLoop() {
	//init golds
	golds = make([]*gold, boardWidth)
	for i := 0; i < boardWidth; i++ {
		g := new(gold)
		golds[i] = g
		g.c = termbox.ColorYellow
		g.score = 100
		g.y = rand.Intn(boardHeight + boardHeight)
		if g.y > boardHeight {
			g.cd = g.y
			g.y = -1
		}
	}
	tick := time.NewTicker(300 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			for i, g := range golds {
				if g.cd > 0 {
					g.cd--
					if g.cd > 0 {
						continue
					}
				}
				if g.y >= 0 {
					termbox.SetCell(i, g.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
				}
				g.y++
				if g.y >= boardHeight {
					if i == px {
						point += g.score
						refreshScore()
					}
					g.y = -1
					g.cd = rand.Intn(boardHeight << 2)
					continue
				}
				termbox.SetCell(i, g.y, ' ', termbox.ColorYellow, termbox.ColorYellow)
			}
			termbox.Flush()
		}
	}
}
func movePlayer(delta int) {
	if px == 0 && delta < 0 || px >= boardWidth && delta > 0 {
		return
	}
	termbox.SetCell(px, boardHeight, ' ', termbox.ColorDefault, termbox.ColorDefault)
	px += delta
	termbox.SetCell(px, boardHeight, '_', termbox.ColorCyan, termbox.ColorDefault)
	termbox.Flush()
}

func refreshScore() {
	x := boardWidth + 10
	p := point
	if p == 0 {
		for i := 0; i < 3; i++ {
			termbox.SetCell(x, 3, '0', termbox.ColorGreen, termbox.ColorWhite)
			x--
		}
	}
	for p > 0 {
		termbox.SetCell(x, 3, []rune(fmt.Sprintf("%d", p%10))[0], termbox.ColorGreen, termbox.ColorWhite)
		p /= 10
		x--
	}
	termbox.Flush()
}

func main() {
	//termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	px = boardWidth >> 1
	termbox.SetCell(px, boardHeight, '_', termbox.ColorCyan, termbox.ColorDefault)
	termbox.Flush()
	refreshScore()
	go goldRainLoop()
	//input
	termbox.SetInputMode(termbox.InputEsc)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Ch == 'Q' || ev.Key == termbox.KeyEsc {
				break mainloop
			} else if ev.Key == termbox.KeyArrowLeft {
				movePlayer(-1)
			} else if ev.Key == termbox.KeyArrowRight {
				movePlayer(1)
			}
		default:
		}
	}
}
