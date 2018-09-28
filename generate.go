package main

import (
    "github.com/faiface/pixel"
    "math/rand"
    "time"
    "image/color"
    "golang.org/x/image/colornames"
)

//Индекс начала генерации
var indexOfBeginGenerate = "1|1"
//Ширина кубов влабиринта
var factor = 10
//Ширина лабиринта
var lengthX = 101
//Высота либиринта
var lengthY = 77
//Лист перекрестков(для генерации)
var crossways []string
var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

//Индексы лабиринта
var collection = map[string]Square{}

const WallColor = 0
const ChannelColor = 1
const PlayerColor = 2
const FinishColor = 3
const PlayerWayColor = 4
const BotColor = 5
const BotWayColor = 6

//Цвета либинта
var colorMap = map[int]color.RGBA{
    WallColor:      colornames.Black,
    ChannelColor:   colornames.White,
    PlayerColor:    colornames.Blue,
    FinishColor:    colornames.White,
    PlayerWayColor: colornames.Lightblue,
    BotColor:       colornames.Red,
    BotWayColor:    colornames.Lightsalmon,
}
//Куб
type Square struct {
    state int
    color color.RGBA
    rect  pixel.Rect
    x     int
    y     int
    pass  bool
}

//назначение кода цвета
func (s *Square) getColor() color.RGBA {
    return colorMap[s.state]
}

//Генерация матрицы лабиринта
func generate(lengthX int, lengthY int) {
    //Create blank
    collection = make(map[string]Square)
    for y := 0; y < lengthY; y++ {
        for x := 0; x < lengthX; x++ {
            rect := pixel.R(float64(x*factor), float64(y*factor), float64(factor+x*factor), float64(factor+y*factor))
            square := Square{x: x, y: y, rect: rect}
            if x+1 == lengthX || x == 0 || y+1 == lengthY || y == 0 || x%2 != 1 || y%2 != 1 {
                square.state = WallColor
            } else {
                square.state = ChannelColor
            }
            collection[getIndex(x, y)] = square
        }
    }
    //Create paths
    updateWay(indexOfBeginGenerate)

    for len(crossways) > 0 {
        crossway := crossways[0]
        crossways = crossways[1:]
        updateWay(crossway)
    }

    for _, item := range collection {
        item.pass = false
        item.save()
    }

    return
}

func (s Square) save() {
    collection[getIndex(s.x, s.y)] = s
}

//Делает одну итерацию генерации(генерирует ветку маршрута)
func updateWay(indexWay string) {
    siblings := collection[indexWay].getSiblings(2)
    if len(siblings) > 1 {
        crossways = append(crossways, indexWay)
    }
    for len(siblings) > 0 {
        if len(siblings) > 1 {
            crossways = append(crossways, indexWay)
        }
        if len(siblings) > 0 {
            nextIndex := siblings[r.Intn(len(siblings))]
            for _, index := range getMediator(collection[indexWay], collection[nextIndex]) {
                element := collection[index]
                element.state = ChannelColor
                element.pass = true
                element.save()
                indexWay = index
            }
        }
        siblings = collection[indexWay].getSiblings(2)
    }
}

//получить соседей
func (s Square) getSiblings(step int) (list []string) {
    if current, ok := collection[getIndex(s.x, s.y-step)]; ok && current.pass == false && current.state != 0 {
        list = append(list, getIndex(s.x, s.y-step))
    }
    if current, ok := collection[getIndex(s.x, s.y+step)]; ok && current.pass == false && current.state != 0 {
        list = append(list, getIndex(s.x, s.y+step))
    }
    if current, ok := collection[getIndex(s.x-step, s.y)]; ok && current.pass == false && current.state != 0 {
        list = append(list, getIndex(s.x-step, s.y))
    }
    if current, ok := collection[getIndex(s.x+step, s.y)]; ok && current.pass == false && current.state != 0 {
        list = append(list, getIndex(s.x+step, s.y))
    }
    return
}
//отдает все индексы от одного куба до другого
func getMediator(square1 Square, square2 Square) (list []string) {
    for square1.x > square2.x {
        square1.x--
        list = append(list, getIndex(square1.x, square1.y))
    }

    for square1.x < square2.x {
        square1.x++
        list = append(list, getIndex(square1.x, square1.y))
    }

    for square1.y > square2.y {
        square1.y--
        list = append(list, getIndex(square1.x, square1.y))
    }

    for square1.y < square2.y {
        square1.y++
        list = append(list, getIndex(square1.x, square1.y))
    }
    return
}
