package app_test

import (
	"github.com/andygeiss/ecs-example/internal/app"
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestGenerateEntities(t *testing.T) {
	out := app.GenerateEntities(1000, 1920, 1080)
	assert.That("entitiy count should be 1000", t, len(out), 1000)
}
