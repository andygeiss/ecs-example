package entities

import (
	"fmt"
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/utils/engine/entity"
	"math/rand"
)

// Generate ...
func Generate(num, width, height int) []*entity.Entity {
	out := make([]*entity.Entity, num)
	for i := range out {
		out[i] = entity.NewEntity(fmt.Sprintf("%d", i), []entity.Component{
			components.NewPosition(rand.Float32()*float32(width), rand.Float32()*float32(height)),
			components.NewSize(3, 3),
			components.NewVelocity(rand.Float32()*10, rand.Float32()*10),
		})
	}
	return out
}
