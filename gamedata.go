package main

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	gd := GameData{
		ScreenHeight: 80,
		ScreenWidth:  50,
		TileWidth:    16,
		TileHeight:   16,
	}
	return gd
}
