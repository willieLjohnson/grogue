package main

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	gd := GameData{
		ScreenHeight: 50,
		ScreenWidth:  80,
		TileWidth:    16,
		TileHeight:   16,
	}
	return gd
}
