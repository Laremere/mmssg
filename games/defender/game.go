package defender

import (
	"github.com/laremere/mmssg/games"
	"log"
)

func init() {
	game.Register("defender", NewGame)
}

func NewGame() game.Game {
	return new(defender)
}

type defender struct {
}

func (d *defender) UserEvent(event game.UserEvent) {
	switch event := event.(type) {
	case *game.UserMoveEvent:
	case *game.UserAddEvent:
		log.Println("defender user added: ", event)
	case *game.UserDropEvent:
	default:
		log.Fatal("Unhandled user event type!")
	}
}
func (d *defender) Update() string { return "" }
