package main

import (
	"github.com/andygeiss/ecs-example/internal/app"
	"github.com/andygeiss/utils/engine/entity"
	"github.com/andygeiss/utils/run"
)

func main() {
	cfg := app.Config{
		NumberOfEntities: 200000,
		Width:            1366,
		Height:           768,
		Title:            "ECS Example Benchmark",
	}
	run.Main(func() {
		app.Entrypoint(&cfg, entity.DefaultRepository)
	})
}
