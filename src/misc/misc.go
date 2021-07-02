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

func StringToPlayer(str string) Player {
	for i, val := range playerMapping {
		if val == str {
			return Player(i)
		}
	}
	return -1
}

func StringToIntPlayer(str string) int {
	return int(StringToPlayer(str))
}
