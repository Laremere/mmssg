package game

type NewGame func() Game

type Game interface {
	UserEvent(event UserEvent)
	Update() string
}

func Register(name string, initialize NewGame) {
	if Games == nil {
		Games = make(map[string]NewGame)
	}
	Games[name] = initialize
}

var Games map[string]NewGame

type UserEvent interface{}

type UserAddEvent struct {
	Id int
}

type UserDropEvent struct {
	Id int
}

type UserMoveEvent struct {
	Id   int
	X, Y float64
}

type UserActionEvent struct {
	Id int
}
