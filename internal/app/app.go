package app

import (
	"github.com/andygeiss/ecs-example/internal/plugins"
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/engines"
	"github.com/andygeiss/ecs/entities"
	"github.com/andygeiss/ecs/systems"
)

// App ...
type App struct {
	engine      core.Engine
	numEntities int
	width       float32
	height      float32
	title       string
}

// Setup ...
func (a *App) Setup() {
	em := entities.NewEntityManager()
	em.Add(GenerateEntities(a)...)
	sm := systems.NewSystemManager()
	sm.Add(
		NewMovement(),
		NewCollision(a.width, a.height),
		NewRendering().(*Rendering).
			WithWidth(a.width).
			WithHeight(a.height).
			WithTitle(a.title).
			WithPlugins(plugins.ShowEngineStats(em)),
	)
	engine := engines.NewDefaultEngine(em, sm)
	engine.Setup()
	a.engine = engine
}

// Run ...
func (a *App) Run() {
	a.engine.Run()
}

// Teardown ...
func (a *App) Teardown() {
	a.engine.Teardown()
}

func (a *App) WithNumEntities(numEntities int) *App {
	a.numEntities = numEntities
	return a
}

func (a *App) WithWidth(width float32) *App {
	a.width = width
	return a
}

func (a *App) WithHeight(height float32) *App {
	a.height = height
	return a
}

func (a *App) WithTitle(title string) *App {
	a.title = title
	return a
}

// NewApp ...
func NewApp() core.Engine {
	return &App{}
}
