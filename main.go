package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"strconv"
	"math/rand"
	"fmt"
	"time"
)

var factor = 10
var collection = map[string]Square{}
var beginIndex string

type Square struct {
	state int
	color color.RGBA
	rect  pixel.Rect
	x     int
	y     int
}

func getIndex(x int, y int) string {
	return strconv.Itoa(x) + "|" + strconv.Itoa(y)
}

func generatePattern(lengthX int, lengthY int) {
	collection = make(map[string]Square)
	for y := 0; y < lengthY; y++ {
		for x := 0; x < lengthX; x++ {
			rect := pixel.R(float64(x*factor), float64(y*factor), float64(factor+x*factor), float64(factor+y*factor))
			square := Square{x: x, y: y, rect: rect}
			if x+1 == lengthX || x == 0 || y+1 == lengthY || y == 0 || x%2 != 1 || y%2 != 1 {
				square.state = 0
			} else {
				if beginIndex != "" && rand.Intn(2) == 0 {
					beginIndex = getIndex(x, y)
				}
				square.state = 1
			}
			collection[getIndex(x, y)] = square
		}
	}

	return
}

func getSiblings() (list []string) {
	rect := collection[beginIndex]
	if _, ok := collection[getIndex(rect.x, rect.y-2)]; ok {
		list = append(list, getIndex(rect.x, rect.y-2))
	}
	if _, ok := collection[getIndex(rect.x, rect.y+2)]; ok {
		list = append(list, getIndex(rect.x, rect.y+2))
	}
	if _, ok := collection[getIndex(rect.x-2, rect.y)]; ok {
		list = append(list, getIndex(rect.x-2, rect.y))
	}
	if _, ok := collection[getIndex(rect.x+2, rect.y)]; ok {
		list = append(list, getIndex(rect.x+2, rect.y))
	}
	return
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1010, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(win)
	for _, square := range collection {
		imd.Color = colornames.Black
		imd.Push(square.rect.Min, square.rect.Max)
		imd.Rectangle(float64(square.state))
	}

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	generatePattern(101, 77)
	checkPointBegin()
	pixelgl.Run(run)
}

func randBool() bool {
	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)
	return r.Intn(2) != 0

}

func checkPointBegin() {
	r := rand.New(rand.NewSource(99))
	fmt.Println(r.Int())
}
