package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var window *pixelgl.Window

func run() {
	var win, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Random Labyrinth!",
		Bounds: pixel.R(0, 0, 1010, 770),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}
	window = win
	for !window.Closed() {
		imd := imdraw.New(window)
		window.Clear(colornames.White)
		keywordEvents()
		bot.autoSearch()
		for _, square := range collection {
			imd.Color = square.getColor()
			imd.Push(square.rect.Min, square.rect.Max)
			imd.Rectangle(0)
		}
		imd.Draw(window)
		window.Update()
	}
}

func main() {
	generate(lengthX, lengthY)
	initPlayers()
	pixelgl.Run(run)
}

func keywordEvents() {
	destinationIndex := getPlayerRouteIndex()
	if destinationIndex != player.currentIndex {
		for _, sibling := range player.getPlayerIndex().getSiblings(1, []int{}, false) {
			if sibling == destinationIndex {
				player.doStep(destinationIndex)
			}
		}
	}
}

func initPlayers() {
	player.setStart()
	player.setFinish()
	bot = player
	bot.color = BotColor
	bot.colorWay = BotWayColor
}
