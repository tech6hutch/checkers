package pieces

type Color int

const (
	ColorDark Color = iota
	ColorLight
)

func (color Color) String() string {
	switch color {
	case ColorDark: return "dark"
	case ColorLight: return "light"
	default: panic("unknown variant")
	}
}

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
