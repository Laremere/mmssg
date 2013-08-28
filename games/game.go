package game

type NewGame func() Game

type Game interface {
	UserAdd(id int)
	UserDrop(id int)
	UserUpdate(id int, vx, vy float64)
	UserAction(id int)
	Update() string
}

func Register(name string, initialize NewGame) {
	if Games == nil {
		Games = make(map[string]NewGame)
	}
	Games[name] = initialize
}

var Games map[string]NewGame
