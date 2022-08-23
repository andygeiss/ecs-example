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
	em.Add(GenerateEntities(a.numEntities, a.width, a.height)...)
	sm := systems.NewSystemManager()
	sm.Add(NewMovement(), NewCollision(a.width, a.height), NewRendering(int32(a.width), int32(a.height), "", plugins.ShowEngineStats(em)))
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

// NewApp ...
func NewApp(numEntities int, width, height float32, title string) core.Engine {
	return &App{
		numEntities: numEntities,
		width:       width,
		height:      height,
		title:       title,
	}
}
