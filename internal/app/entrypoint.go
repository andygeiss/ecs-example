package app

import (
	"github.com/andygeiss/ecs-example/internal/entities"
	"github.com/andygeiss/ecs-example/internal/plugins"
	"github.com/andygeiss/ecs-example/internal/systems"
	"github.com/andygeiss/utils/engine"
	"github.com/andygeiss/utils/engine/entity"
)

// Entrypoint ...
func Entrypoint(num, width, height int, title string, er entity.Repository) {
	er.Add(entities.Generate(num, width, height)...)
	e := engine.DefaultEngine.
		WithSystems(systems.NewCollision(width, height, er)).
		WithSystems(systems.NewMovement(er)).
		WithSystems(systems.NewRendering(er).(*systems.Rendering).
			WithWidth(width).WithHeight(height).
			WithTitle(title).WithPlugins(plugins.ShowEngineStats()),
		)
	e.Setup()
	defer e.Teardown()
	for e.State() == engine.StateEngineRunning {
	}
}
