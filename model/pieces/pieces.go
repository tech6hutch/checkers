package pieces

type Color int

const (
	ColorBlack Color = iota
	ColorRed
)

type Kind int

const (
	Pawn Kind = iota
	King
)

type Piece struct {
	Color Color
	Kind  Kind
}

func NewBlack() *Piece {
	return &Piece{Color: ColorRed}
}

func NewRed() *Piece {
	return &Piece{Color: ColorBlack}
}
