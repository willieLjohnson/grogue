package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

type Game struct {
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, tags := InitWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	return g
}

func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Grogue")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
