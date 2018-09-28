package main

import (
    "github.com/faiface/pixel/pixelgl"
)

var player = Player{color: PlayerColor, colorWay: PlayerWayColor}
var bot = Player{color: BotColor, colorWay: BotWayColor}

type Player struct {
    currentIndex string
    startIndex   string
    finishIndex  string
    color        int
    colorWay     int
    history      []string
}

func getPlayerRouteIndex() (index string) {
    switch true {
    case window.Pressed(pixelgl.KeyLeft):
        return getIndex(player.getPlayerIndex().x-1, player.getPlayerIndex().y)
    case window.Pressed(pixelgl.KeyRight):
        return getIndex(player.getPlayerIndex().x+1, player.getPlayerIndex().y)
    case window.Pressed(pixelgl.KeyDown):
        return getIndex(player.getPlayerIndex().x, player.getPlayerIndex().y-1)
    case window.Pressed(pixelgl.KeyUp):
        return getIndex(player.getPlayerIndex().x, player.getPlayerIndex().y+1)
    }
    return player.currentIndex
}

func doStep(destinationIndex string) {
    c := player.getPlayerIndex()
    c.state = PlayerWayColor
    c.save()
    player.currentIndex = destinationIndex
    c = player.getPlayerIndex()
    c.state = PlayerColor
    c.save()
}

func (player *Player) getPlayerIndex() *Square {
    s := collection[player.currentIndex]
    return &s
}

//выбор точки старта
func (player *Player) setStart() {
    var x, y int
    if r.Intn(2) == 1 {
        x = oddRandom(lengthX)
        y = 0
    } else {
        x = 0
        y = oddRandom(lengthY)
    }
    player.startIndex = getIndex(x, y)
    player.currentIndex = player.startIndex
    element := collection[player.startIndex]
    element.state = PlayerColor
    element.pass = true
    element.save()
}

//Выбор точки финиша
func (player *Player) setFinish() {
    element := collection[player.startIndex]
    x := lengthX - (element.x + 1)
    y := lengthY - (element.y + 1)
    player.finishIndex = getIndex(x, y)
    element = collection[player.finishIndex]
    element.state = FinishColor
    element.pass = true
    element.save()
}

func (player *Player) autoSearch() {

    siblings := collection[player.currentIndex].getSiblings(1)
    if len(siblings) > 1 {
        player.history = append(player.history, player.currentIndex)
    }
    //player.autoSearch()
}
