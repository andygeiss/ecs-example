package systems

import (
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/utils/engine"
	"github.com/andygeiss/utils/engine/entity"
)

// Movement ...
type Movement struct {
	err  error
	repo entity.Repository
}

func (a *Movement) Error() (err error) {
	return a.err
}

func (a *Movement) Setup() {}

func (a *Movement) Process(stopCh chan bool) {
	for _, e := range a.repo.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
}

func (a *Movement) Teardown() {

}

// NewMovement ...
func NewMovement(repo entity.Repository) engine.System {
	return &Movement{
		repo: repo,
	}
}
