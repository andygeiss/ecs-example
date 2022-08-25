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
			components.NewPosition(rand.Float32()*float32(cfg.Width), rand.Float32()*float32(cfg.Height)),
			components.NewSize(3, 3),
			components.NewVelocity(rand.Float32()*10, rand.Float32()*10),
		})
	}
	return out
}
