package entities_test

import (
	"github.com/andygeiss/ecs-example/internal/entities"
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	out := entities.Generate(1000, 1024, 768)
	assert.That("entitiy count should be 1000", t, len(out), 1000)
}
