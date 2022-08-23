package app

import (
	"fmt"
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/ecs/core"
	"math/rand"
)

// GenerateEntities ...
func GenerateEntities(num int, width, height float32) []*core.Entity {
	out := make([]*core.Entity, num)
	for i := range out {
		out[i] = core.NewEntity(fmt.Sprintf("%d", i), []core.Component{
			components.NewPosition(rand.Float32()*width, rand.Float32()*height),
			components.NewSize(3, 3),
			components.NewVelocity(rand.Float32()*10, rand.Float32()*10),
		})
	}
	return out
}
