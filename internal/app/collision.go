package app

import (
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/ecs/core"
)

// Collision ...
type Collision struct {
	height float32
	width  float32
}

func (m *Collision) Setup() {}

func (m *Collision) Process(em core.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := entity.Get(components.MaskPosition).(*components.Position)
		velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
		if position.X >= m.width || position.X <= 0 {
			velocity.X = -velocity.X
		}
		if position.Y >= m.height || position.Y <= 0 {
			velocity.Y = -velocity.Y
		}
	}
	return core.StateEngineContinue
}

// NewCollision ...
func NewCollision(width, height float32) core.System {
	return &Collision{
		height: height,
		width:  width,
	}
}

func (m *Collision) Teardown() {}
