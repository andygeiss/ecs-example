package components

const (
	MaskPosition = uint64(1 << 0)
)

// Position ...
type Position struct {
	ID string  `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}

// NewPosition ...
func NewPosition(x, y float32) *Position {
	return &Position{
		ID: "position",
		X:  x,
		Y:  y,
	}
}

func (p *Position) Mask() uint64 {
	return MaskPosition
}
