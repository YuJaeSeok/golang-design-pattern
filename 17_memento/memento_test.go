package memento

import "testing"

func ExampleGame() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	process := game.Save()
	
	game.Play(-2, -3)
	game.Status()

	game.Load(process)
	game.Status()
}

func TestGame(t *testing.T) {
	ExampleGame()
}
