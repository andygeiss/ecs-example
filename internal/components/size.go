package components

const (
	MaskSize = uint64(1 << 1)
)

// Size ...
type Size struct {
	ID     string  `json:"id"`
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
}

// NewSize ...
func NewSize(width, height float32) *Size {
	return &Size{
		ID:     "size",
		Width:  width,
		Height: height,
	}
}

func (s *Size) Mask() uint64 {
	return MaskSize
}
