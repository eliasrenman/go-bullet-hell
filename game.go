package main

type Game struct {
	player Player
}

func InitalizeGame() *Game {
	game := Game{
		player: InitalizePlayer(),
	}
	return &game
}
