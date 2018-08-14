package main

import (
    "golang.org/x/image/colornames"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel/imdraw"
)

var window *pixelgl.Window
var bot Bot

type Bot struct {
    currentIndex string
}

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
        bot.botProcess()
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
    player = Player{currentIndex: currentIndex}
    bot = Bot{currentIndex: currentIndex}
    pixelgl.Run(run)
}

func keywordEvents() {
    destinationIndex := getPlayerRouteIndex()
    if destinationIndex != player.currentIndex {
        for _, sibling := range player.getPlayerIndex().getSiblings(1) {
            if sibling == destinationIndex {
                doStep(destinationIndex)
            }
        }
    }
}

func (bot Bot) botProcess(){
    bot.getPoint().getSiblings(1)
}

func (bot Bot) getPoint() *Square {
    s := collection[bot.currentIndex]
    return &s
}