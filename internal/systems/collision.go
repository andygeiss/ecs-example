package systems

import (
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/utils/engine"
	"github.com/andygeiss/utils/engine/entity"
)

// Collision ...
type Collision struct {
	err    error
	height float32
	width  float32
	er     entity.Repository
}

func (a *Collision) Error() error {
	return a.err
}

func (a *Collision) Setup() {}

func (a *Collision) Process(stopCh chan bool) {
	for _, e := range a.er.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		if position.X >= a.width || position.X <= 0 {
			velocity.X = -velocity.X
		}
		if position.Y >= a.height || position.Y <= 0 {
			velocity.Y = -velocity.Y
		}
	}
}

func (a *Collision) Teardown() {

}

// NewCollision ...
func NewCollision(width, height int, er entity.Repository) engine.System {
	return &Collision{
		height: float32(height),
		width:  float32(width),
		er:     er,
	}
}
