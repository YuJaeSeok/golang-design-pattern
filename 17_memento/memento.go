package memento

import "fmt"

type Memento interface{}

type Game struct {
	hp, mp int
}

type GameMemento struct {
	hp, mp int
}

func (g *Game) Play(meDelta, hpDelta int) {
	g.mp += meDelta
	g.hp += hpDelta
}

func (g *Game) Save() Memento {
	return &GameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*GameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}
