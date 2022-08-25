package systems

import (
	"github.com/andygeiss/ecs-example/internal/components"
	"github.com/andygeiss/utils/engine"
	"github.com/andygeiss/utils/engine/entity"
	"github.com/andygeiss/utils/run"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Rendering ...
type Rendering struct {
	err           error
	title         string
	width, height int32
	repo          entity.Repository
	plugins       []engine.Plugin
}

func (a *Rendering) Error() error {
	return a.err
}

func (a *Rendering) Setup() {
	run.Safe(func() {
		rl.InitWindow(a.width, a.height, a.title)
	})
}

func (a *Rendering) Process(stopCh chan bool) {
	shouldStop := false
	run.Safe(func() {
		// First check if app should stop.
		if rl.WindowShouldClose() {
			shouldStop = true
		}
		// Clear the screen
		if rl.IsWindowReady() {
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)
			a.renderEntities()
			rl.DrawFPS(10, 10)
			rl.EndDrawing()
		}
	})
	// check for a stop.
	if shouldStop {
		stopCh <- true
		return
	}
	// Dispatch work to plugins.
	for _, plugin := range a.plugins {
		plugin(a.repo)
	}
}

func (a *Rendering) Teardown() {
	run.Safe(func() {
		rl.CloseWindow()
	})
}

func (a *Rendering) WithHeight(height int) *Rendering {
	a.height = int32(height)
	return a
}

func (a *Rendering) WithPlugins(plugins ...engine.Plugin) *Rendering {
	a.plugins = plugins
	return a
}

func (a *Rendering) WithTitle(title string) *Rendering {
	a.title = title
	return a
}

func (a *Rendering) WithWidth(width int) *Rendering {
	a.width = int32(width)
	return a
}

func (a *Rendering) renderEntities() {
	for _, e := range a.repo.FilterByMask(components.MaskPosition | components.MaskSize) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		rl.DrawRectangleRec(rl.Rectangle{X: position.X, Y: position.Y, Width: size.Width, Height: size.Height}, rl.Blue)
	}
}

// NewRendering ...
func NewRendering(repo entity.Repository) engine.System {
	return &Rendering{
		repo: repo,
	}
}
