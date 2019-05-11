package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	boardWidth   = 50
	boardHeight  = 15
	sceneWelcome = 1
	sceneGame    = 2
)

var (
	scene int
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

//print string in the console
func printTb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

//move gold down from the sky
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
	tick := time.NewTicker(200 * time.Millisecond)
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
					g.cd = rand.Intn(boardHeight * 5)
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

//when the point is changed, refresh scoreboard
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

func welcomeScene() {
	scene = sceneWelcome
	printTb(15, 3, termbox.ColorYellow, termbox.ColorDefault, "Gold coin game")
	printTb(8, 5, termbox.ColorBlue, termbox.ColorDefault, "yellow blocks are gold")
	printTb(8, 6, termbox.ColorBlue, termbox.ColorDefault, "silver underscore is your plate")
	printTb(8, 7, termbox.ColorBlue, termbox.ColorDefault, "press <- and -> to move the plate")
	printTb(8, 8, termbox.ColorBlue, termbox.ColorDefault, "scores gain by catch gold with plate")
	printTb(8, 9, termbox.ColorBlue, termbox.ColorDefault, "press 'q' or [ESC] to quit")
	printTb(8, 10, termbox.ColorBlue, termbox.ColorDefault, "press any key to start game")
	termbox.Flush()
}

func startGame() {
	scene = sceneGame
	px = boardWidth >> 1
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(px, boardHeight, '_', termbox.ColorCyan, termbox.ColorDefault)
	refreshScore()
	go goldRainLoop()
	termbox.Flush()
}

func main() {
	//termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	//input
	termbox.SetInputMode(termbox.InputEsc)
	welcomeScene()
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Ch == 'Q' || ev.Key == termbox.KeyEsc {
				break mainloop
			} else if scene == sceneWelcome {
				startGame()
			} else if ev.Key == termbox.KeyArrowLeft {
				movePlayer(-1)
			} else if ev.Key == termbox.KeyArrowRight {
				movePlayer(1)
			}
		default:
		}
	}
}
