package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

type Level struct {
	Tiles []MapTile
	Rooms []Rect
}

func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	return l
}

func (level *Level) CreateRoom(rect Rect) {
	for y := rect.Y1 + 1; y < rect.Y2; y++ {
		for x := rect.X1 + 1; x < rect.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	tiles := level.CreateTiles()
	level.Tiles = tiles

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth-w-1) - 1
		y := GetDiceRoll(gd.ScreenHeight-h-1) - 1

		newRoom := NewRect(x, y, w, h)
		okToAdd := true
		for _, otherRoom := range level.Rooms {
			if newRoom.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.CreateRoom(newRoom)
			level.Rooms = append(level.Rooms, newRoom)
		}
	}
}
func (level *Level) DrawLevel(screen *ebiten.Image) {
	level.DrawMap(screen)
}

func (level *Level) GetIndexFromXY(x int, y int) int {
	return (y * NewGameData().ScreenWidth) + x
}

func (level *Level) CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:  x * gd.TileWidth,
				PixelY:  y * gd.TileHeight,
				Blocked: true,
				Image:   wall,
			}
			tiles[index] = tile
		}
	}

	return tiles
}

func (level *Level) DrawMap(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, opts)
		}
	}
}
