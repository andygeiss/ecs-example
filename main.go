package main

import (
	"github.com/andygeiss/ecs-example/internal/entities"
	"github.com/andygeiss/ecs-example/internal/plugins"
	"github.com/andygeiss/ecs-example/internal/systems"
	"github.com/andygeiss/utils/engine"
	"github.com/andygeiss/utils/engine/entity"
	"github.com/andygeiss/utils/run"
)

func entrypoint() {
	width, height := 1024, 768
	em := entity.DefaultManager
	em.Add(entities.Generate(200000, width, height)...)
	e := engine.DefaultEngine.
		WithSystems(systems.NewCollision(width, height, em)).
		WithSystems(systems.NewMovement(em)).
		WithSystems(systems.NewRendering(em).(*systems.Rendering).
			WithWidth(width).WithHeight(height).
			WithTitle("").WithPlugins(plugins.ShowEngineStats(em)))
	e.Setup()
	defer e.Teardown()
	for e.State() == engine.StateEngineRunning {
	}
}

func main() {
	run.Main(entrypoint)
}
