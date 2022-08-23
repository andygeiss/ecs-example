package app

import (
	"fmt"
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/ecs/core"
	"math/rand"
)

// GenerateEntities ...
func GenerateEntities(app *App) []*core.Entity {
	out := make([]*core.Entity, app.numEntities)
	for i := range out {
		out[i] = core.NewEntity(fmt.Sprintf("%d", i), []core.Component{
			components.NewPosition(rand.Float32()*app.width, rand.Float32()*app.height),
			components.NewSize(3, 3),
			components.NewVelocity(rand.Float32()*10, rand.Float32()*10),
		})
	}
	return out
}
