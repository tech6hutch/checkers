package pieces

type Color int

const (
	ColorDark Color = iota
	ColorLight
)

type Kind int

const (
	Man Kind = iota
	King
)

type Piece struct {
	Color Color
	Kind  Kind
}

func New(color Color) *Piece {
	return &Piece{Color: color}
}
