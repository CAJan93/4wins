package misc

type Player int

type Direction int

const (
	Horizontal    Direction = iota
	Vertical                = iota
	Diagonal                = iota
	NoPlayerValue           = " "
)

var playerMapping = [...]string{"X", "O"}

// PlayerToString provides mapping for player to string
func PlayerToString(p Player) string {
	return playerMapping[p]
}

// playerIntToStrig is a wrapper for PlayerToString
func PlayerIntToStrig(p int) string {
	return PlayerToString(Player(p))
}
