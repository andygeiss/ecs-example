package components

const (
	MaskVelocity = uint64(1 << 2)
)

// Velocity ...
type Velocity struct {
	ID string  `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}

// NewVelocity ...
func NewVelocity(x, y float32) *Velocity {
	return &Velocity{
		ID: "velocity",
		X:  x,
		Y:  y,
	}
}

func (s *Velocity) Mask() uint64 {
	return MaskVelocity
}
