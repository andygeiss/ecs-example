package main

import (
	"fmt"
	"github.com/andygeiss/ecs-example/engine"
	"github.com/andygeiss/ecs-example/plugins"
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/engines"
	"github.com/andygeiss/ecs/entities"
	"github.com/andygeiss/ecs/systems"
	"github.com/andygeiss/utils/run"
	"math/rand"
)

const (
	Width  = 800
	Height = 600
)

func generateEntities(num int) []*core.Entity {
	out := make([]*core.Entity, num)
	for i := range out {
		out[i] = core.NewEntity(fmt.Sprintf("%d", i), []core.Component{
			engine.NewPosition(rand.Float32()*Width, rand.Float32()*Height),
			engine.NewSize(3, 3),
			engine.NewVelocity(rand.Float32()*10, rand.Float32()*10),
		})
	}
	return out
}

func main() {
	run.Main(func() {
		em := entities.NewEntityManager()
		em.Add(generateEntities(100000)...)
		sm := systems.NewSystemManager()
		sm.Add(
			engine.NewMovement(),
			engine.NewCollision(Width, Height),
			engine.NewRendering(Width, Height, "ECS with Raylib Demo",
				plugins.ShowEngineStats(em),
			),
		)
		e := engines.NewDefaultEngine(em, sm)
		e.Setup()
		defer e.Teardown()
		e.Run()
	})
}
