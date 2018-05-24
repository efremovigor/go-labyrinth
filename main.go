package main

import (
    "golang.org/x/image/colornames"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel/imdraw"
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
    pixelgl.Run(run)

}

func getRouteIndex() (index string) {
    switch true {
    case window.Pressed(pixelgl.KeyLeft):
        return getIndex(getCurrent().x-1, getCurrent().y)
    case window.Pressed(pixelgl.KeyRight):
        return getIndex(getCurrent().x+1, getCurrent().y)
    case window.Pressed(pixelgl.KeyDown):
        return getIndex(getCurrent().x, getCurrent().y-1)
    case window.Pressed(pixelgl.KeyUp):
        return getIndex(getCurrent().x, getCurrent().y+1)
    }
    return currentIndex
}

func keywordEvents() {
    destinationIndex := getRouteIndex()
    if destinationIndex != currentIndex {
        for _, sibling := range getCurrent().getSiblings(1) {
            if sibling == destinationIndex {
                doStep(destinationIndex)
            }
        }
    }
}

func doStep(destinationIndex string) {
    c := getCurrent()
    c.state = PlayerWay
    c.save()
    currentIndex = destinationIndex
    c = getCurrent()
    c.state = Player
    c.save()
}
