package app

import (
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/utils/run"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Rendering ...
type Rendering struct {
	plugins       []core.Plugin
	title         string
	width, height int32
}

func (r *Rendering) Setup() {
	run.Safe(func() {
		rl.InitWindow(r.width, r.height, r.title)
	})
}

func (r *Rendering) Process(em core.EntityManager) (state int) {
	shouldStop := false
	run.Safe(func() {
		// First check if app should stop.
		if rl.WindowShouldClose() {
			shouldStop = true
		}
		// Clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		r.renderEntities(em)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	})
	// check for a stop.
	if shouldStop {
		return core.StateEngineStop
	}
	// Dispatch work to plugins.
	for _, plugin := range r.plugins {
		plugin(em)
	}
	return core.StateEngineContinue
}

func (r *Rendering) Teardown() {
	run.Safe(func() {
		rl.CloseWindow()
	})
}

func (r *Rendering) WithHeight(height float32) *Rendering {
	r.height = int32(height)
	return r
}

func (r *Rendering) WithPlugins(plugins ...core.Plugin) *Rendering {
	r.plugins = plugins
	return r
}

func (r *Rendering) WithTitle(title string) *Rendering {
	r.title = title
	return r
}

func (r *Rendering) WithWidth(width float32) *Rendering {
	r.width = int32(width)
	return r
}

func (r *Rendering) renderEntities(em core.EntityManager) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskSize) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		rl.DrawRectangleRec(rl.Rectangle{X: position.X, Y: position.Y, Width: size.Width, Height: size.Height}, rl.Blue)
	}
}

// NewRendering ...
func NewRendering() core.System {
	return &Rendering{}
}
