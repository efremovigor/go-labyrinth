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
//Индексы начала и конца лабиринта
var startIndex, finishIndex, currentIndex string

const Wall = 0
const Channel = 1
const Player = 2
const Finish = 3
const PlayerWay = 4

//Цвета либинта
var colorMap = map[int]color.RGBA{
    Wall:      colornames.Black,
    Channel:   colornames.White,
    Player:    colornames.Blue,
    Finish:    colornames.White,
    PlayerWay: colornames.Lightblue,
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
                square.state = Wall
            } else {
                square.state = Channel
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
    setStart()
    setFinish()
    return
}

//выбор точки старта
func setStart() {
    var x, y int
    if r.Intn(2) == 1 {
        x = oddRandom(lengthX)
        y = 0
    } else {
        x = 0
        y = oddRandom(lengthY)
    }
    startIndex = getIndex(x, y)
    currentIndex = startIndex
    element := collection[startIndex]
    element.state = Player
    element.pass = true
    element.save()
}

//Выбор точки финиша
func setFinish() {
    element := collection[startIndex]
    x := lengthX - (element.x + 1)
    y := lengthY - (element.y + 1)
    finishIndex = getIndex(x, y)
    element = collection[finishIndex]
    element.state = Finish
    element.pass = true
    element.save()
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
                element.state = Channel
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

func getCurrent() *Square {
    s := collection[currentIndex]
    return &s
}
