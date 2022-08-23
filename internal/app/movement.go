package app

import (
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/ecs/core"
)

// Movement ...
type Movement struct{}

func (m *Movement) Setup() {}

func (m *Movement) Process(em core.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := entity.Get(components.MaskPosition).(*components.Position)
		velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
	return core.StateEngineContinue
}

func (m *Movement) Teardown() {

}

// NewMovement ...
func NewMovement() core.System {
	return &Movement{}
}
