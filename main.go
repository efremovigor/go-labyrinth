package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"strconv"
	"math/rand"
	"time"
	"image/color"
	"fmt"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

var factor = 10
var lengthX = 99
var lengthY = 71
var collection = map[string]Square{}
var beginIndex = "1|1"
var currentIndex string
var crossways []string
var colorMap = map[int]color.RGBA{
	0: colornames.Black,
	1: colornames.White,
	2: colornames.Aqua,
	11: colornames.Red,
}

type Square struct {
	state int
	color color.RGBA
	rect  pixel.Rect
	x     int
	y     int
	pass  bool
}

func (s *Square) getColor() color.RGBA {
	return colorMap[s.state]
}

func getIndex(x int, y int) string {
	return strconv.Itoa(x) + "|" + strconv.Itoa(y)
}

func generate(lengthX int, lengthY int) {
	//Create blank
	collection = make(map[string]Square)
	for y := 0; y < lengthY; y++ {
		for x := 0; x < lengthX; x++ {
			rect := pixel.R(float64(x*factor), float64(y*factor), float64(factor+x*factor), float64(factor+y*factor))
			square := Square{x: x, y: y, rect: rect}
			if x+1 == lengthX || x == 0 || y+1 == lengthY || y == 0 || x%2 != 1 || y%2 != 1 {
				square.state = 0
			} else {
				if beginIndex == "" && randBool((lengthX+lengthY)/2) {
					beginIndex = getIndex(x, y)
				} else {
					square.state = 1
				}
			}
			collection[getIndex(x, y)] = square
		}
	}
	//Create paths
	currentIndex = beginIndex
	updateWay(currentIndex)

	for len(crossways) > 0 {
		crossway := crossways[0]
		crossways = crossways[1:]
		updateWay(crossway)

	}
	return
}

func setPass(index string) {
	element := collection[index]
	element.state = 1
	element.pass = true
	collection[index] = element
}

func updateWay(indexWay string){
	siblings := getSiblings(indexWay)
	if len(siblings) > 1 {
		crossways = append(crossways,currentIndex)
	}
	for len(siblings) > 0 {
		if len(siblings) > 1 {
			crossways = append(crossways,indexWay)
		}
		if len(siblings) > 0 {
			nextIndex := siblings[r.Intn(len(siblings))]
			for _, index := range getMediator(collection[indexWay], collection[nextIndex]) {
				setPass(index)
				fmt.Println(index)
				indexWay = index
			}
		}
		siblings = getSiblings(indexWay)
	}
}

func getSiblings(index string) (list []string) {
	rect := collection[index]
	if current, ok := collection[getIndex(rect.x, rect.y-2)]; ok && current.pass == false {
		list = append(list, getIndex(rect.x, rect.y-2))
	}
	if current, ok := collection[getIndex(rect.x, rect.y+2)]; ok && current.pass == false {
		list = append(list, getIndex(rect.x, rect.y+2))
	}
	if current, ok := collection[getIndex(rect.x-2, rect.y)]; ok && current.pass == false {
		list = append(list, getIndex(rect.x-2, rect.y))
	}
	if current, ok := collection[getIndex(rect.x+2, rect.y)]; ok && current.pass == false {
		list = append(list, getIndex(rect.x+2, rect.y))
	}
	return
}

func getMediator(square1 Square, square2 Square) (list []string) {
	for square1.x > square2.x {
		square1.x--
		list = append(list,getIndex(square1.x, square1.y))
	}

	for square1.x < square2.x {
		square1.x++
		list = append(list,getIndex(square1.x, square1.y))
	}

	for square1.y > square2.y {
		square1.y--
		list = append(list,getIndex(square1.x, square1.y))
	}

	for square1.y < square2.y {
		square1.y++
		list = append(list,getIndex(square1.x, square1.y))
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

func randBool(limit int) bool {
	return r.Intn(limit) == 0
}
