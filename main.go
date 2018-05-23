package main

import (
    "golang.org/x/image/colornames"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel/imdraw"
)

func run() {
    cfg := pixelgl.WindowConfig{
        Title:  "Random Labyrinth!",
        Bounds: pixel.R(0, 0, 1010, 768),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }
    imd := imdraw.New(win)
    //if win.Pressed(pixelgl.KeyLeft) {
    //    camPos.X -= camSpeed * dt
    //}
    if win.Pressed(pixelgl.KeyRight) {

    }
    //if win.Pressed(pixelgl.KeyDown) {
    //    camPos.Y -= camSpeed * dt
    //}
    //if win.Pressed(pixelgl.KeyUp) {
    //   camPos.Y += camSpeed * dt
    //}
    for _, square := range collection {
        imd.Color = square.getColor()
        imd.Push(square.rect.Min, square.rect.Max)
        imd.Rectangle(0)
    }

    for !win.Closed() {
        win.Clear(colornames.White)
        imd.Draw(win)
        win.Update()
    }
}

func main() {
    generate(lengthX, lengthY)
    pixelgl.Run(run)

}

func playerStep()  {

}

