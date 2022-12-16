package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var position *ecs.Component
var renderable *ecs.Component

func InitWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{Image: playerImg}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		})
	players := ecs.BuildTag(player, position)
	tags["players"] = players
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables
	return manager, tags
}
