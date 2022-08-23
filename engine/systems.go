package engine

import (
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/utils/run"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
 ____            _
/ ___| _   _ ___| |_ ___ _ __ ___  ___
\___ \| | | / __| __/ _ \ '_ ` _ \/ __|
 ___) | |_| \__ \ ||  __/ | | | | \__ \
|____/ \__, |___/\__\___|_| |_| |_|___/
       |___/
*/

// Collision ...
type Collision struct {
	height float32
	width  float32
}

func (m *Collision) Setup() {}

func (m *Collision) Process(em core.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(MaskPosition | MaskVelocity) {
		position := entity.Get(MaskPosition).(*Position)
		velocity := entity.Get(MaskVelocity).(*Velocity)
		if position.X >= m.width || position.X <= 0 {
			velocity.X = -velocity.X
		}
		if position.Y >= m.height || position.Y <= 0 {
			velocity.Y = -velocity.Y
		}
	}
	return core.StateEngineContinue
}

// NewCollision ...
func NewCollision(width, height float32) core.System {
	return &Collision{
		height: height,
		width:  width,
	}
}

func (m *Collision) Teardown() {}

// Movement ...
type Movement struct{}

func (m *Movement) Setup() {}

func (m *Movement) Process(em core.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(MaskPosition | MaskVelocity) {
		position := entity.Get(MaskPosition).(*Position)
		velocity := entity.Get(MaskVelocity).(*Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
	return core.StateEngineContinue
}

func (m *Movement) Teardown() {

}

// NewMovement ...
func NewMovement() core.System {
	return &Movement{}
}

// Rendering ...
type Rendering struct {
	plugins       []core.Plugin
	title         string
	width, height int32
}

func (r *Rendering) Setup() {
	run.Safe(func() {
		rl.InitWindow(r.width, r.height, r.title)
		rl.SetTargetFPS(60)
	})
}

func (r *Rendering) Process(em core.EntityManager) (state int) {
	shouldStop := false
	run.Safe(func() {
		// First check if engine should stop.
		if rl.WindowShouldClose() {
			shouldStop = true
		}
		// Clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// Render entities
		for _, e := range em.FilterByMask(MaskPosition | MaskSize) {
			position := e.Get(MaskPosition).(*Position)
			size := e.Get(MaskSize).(*Size)
			rl.DrawRectangleRec(rl.Rectangle{X: position.X, Y: position.Y, Width: size.Width, Height: size.Height}, rl.Red)
		}
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

// NewRendering ...
func NewRendering(width, height int32, title string, plugins ...core.Plugin) core.System {
	return &Rendering{
		height:  height,
		plugins: plugins,
		title:   title,
		width:   width,
	}
}
