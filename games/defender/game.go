package defender

import (
	"github.com/laremere/mmssg/games"
)

func init() {
	game.Register("defender", NewGame)
}

func NewGame() game.Game {
	return new(defender)
}

type defender struct {
}

func (d *defender) UserAdd(id int)                    {}
func (d *defender) UserDrop(id int)                   {}
func (d *defender) UserUpdate(id int, vx, vy float64) {}
func (d *defender) UserAction(id int)                 {}
func (d *defender) Update() string                    { return "" }
