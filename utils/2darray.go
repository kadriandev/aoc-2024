package utils

type Point struct {
	x int
	y int
}

type Direction int

const (
	Up Direction = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

var direction = map[Direction]Point{
	Up:        {x: 0, y: 1},
	UpRight:   {x: 1, y: 1},
	Right:     {x: 1, y: 0},
	DownRight: {x: 1, y: -1},
	Down:      {x: 0, y: -1},
	DownLeft:  {x: -1, y: -1},
	Left:      {x: -1, y: 0},
	UpLeft:    {x: -1, y: 1},
}

type TwoDArray struct {
	Board [][]rune
}

func (b TwoDArray) FindWord(word string, x int, y int, d Direction) bool {
	if len(word) == 0 {
		return true
	} else if x < 0 || x >= len(b.Board) || y < 0 || y >= len(b.Board[x]) {
		return false
	}

	runes := []rune(word)
	if b.Board[x][y] == runes[0] {
		return b.FindWord(word[1:], x+direction[d].x, y+direction[d].y, d)
	}
	return false
}
