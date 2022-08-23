package main

import (
	"github.com/andygeiss/ecs-example/internal/app"
	"github.com/andygeiss/utils/run"
)

func main() {
	run.Main(func() {
		app.Entrypoint()
	})
}
