package main

import (
	"github.com/andygeiss/ecs-example/internal/app"
	"github.com/andygeiss/utils/engine/entity"
	"github.com/andygeiss/utils/run"
)

func main() {
	run.Main(func() {
		app.Entrypoint(200000, 1366, 768, "ECS benchmark example", entity.DefaultRepository)
	})
}
