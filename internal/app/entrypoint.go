package app

import (
	"github.com/andygeiss/ecs-example/internal/plugins"
	"github.com/andygeiss/ecs-example/internal/systems"
	"github.com/andygeiss/utils/engine"
	"github.com/andygeiss/utils/engine/entity"
)

// Entrypoint ...
func Entrypoint(cfg *Config, er entity.Repository) {
	er.Add(Generate(cfg)...)
	e := engine.DefaultEngine.
		WithSystems(
			systems.NewCollision(er).(*systems.Collision).
				WithWidth(cfg.Width).
				WithHeight(cfg.Height),
			systems.NewMovement(er),
			systems.NewRendering(er).(*systems.Rendering).
				WithWidth(cfg.Width).
				WithHeight(cfg.Height).
				WithTitle(cfg.Title).
				WithPlugins(plugins.ShowEngineStats()),
		)
	e.Setup()
	defer e.Teardown()
	for e.State() == engine.StateEngineRunning {
	}
}
