package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

type Game struct {
	Map GameMap
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
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
