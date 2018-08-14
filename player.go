package main

import (
    "github.com/faiface/pixel/pixelgl"
)

var player Player

type Player struct {
    currentIndex string
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

func (player Player) getPlayerIndex() *Square {
    s := collection[player.currentIndex]
    return &s
}
