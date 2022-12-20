package main

import "github.com/hajimehoshi/ebiten/v2"

func TryMovePlayer(g *Game) {
	players := g.WorldTags["players"]

	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		y = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		x = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		x = 1
	}

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		tile := level.Tiles[index]
		if tile.Blocked != true {
			pos.X += x
			pos.Y += y
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 16)
		}
	}
	if x != 0 || y != 0 {
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
