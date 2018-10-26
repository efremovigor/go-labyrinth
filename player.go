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

func (player *Player) doStep(destinationIndex string) {
	c := player.getPlayerIndex()
	c.state = player.colorWay
	c.save()
	player.currentIndex = destinationIndex
	c = player.getPlayerIndex()
	c.state = player.color
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
	element.state = player.colorWay
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
	element.pass = false
	element.save()
}

func (player *Player) autoSearch() {
	if player.currentIndex == player.finishIndex {
		return
	}
	siblings := collection[player.currentIndex].getSiblings(1, []int{player.colorWay})
	var next string
	if len(siblings) > 1 {
		player.history = append(player.history, player.currentIndex)
	}

	if len(siblings) == 0 && len(player.history) > 0 {
		current := player.history[len(player.history)-1]
		player.history = player.history[:len(player.history)-1]
		currentSquare := player.getPlayerIndex()
		currentSquare.state = player.colorWay
		currentSquare.save()
		player.currentIndex = current
		player.autoSearch()
	}

	if len(siblings) > 1 {
		next = siblings[r.Intn(len(siblings))]
	} else if len(siblings) == 1 {
		next = siblings[0]
	} else {
		return
	}

	player.doStep(next)
}
