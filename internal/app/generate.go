package app

import (
	"fmt"
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/utils/engine/entity"
	"math/rand"
)

// Generate ...
func Generate(cfg *Config) []*entity.Entity {
	out := make([]*entity.Entity, cfg.NumberOfEntities)
	for i := range out {
		out[i] = entity.NewEntity(fmt.Sprintf("%d", i), []entity.Component{
			components.NewPosition().(*components.Position).
				WithX(rand.Float32() * float32(cfg.Width)).
				WithY(rand.Float32() * float32(cfg.Height)),
			components.NewSize().(*components.Size).
				WithWidth(3).
				WithHeight(3),
			components.NewVelocity().(*components.Velocity).
				WithX(rand.Float32() * 10).
				WithY(rand.Float32() * 10),
		})
	}
	return out
}
