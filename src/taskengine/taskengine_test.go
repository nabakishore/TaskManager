package taskengine

import (
	"testing"
)

func TestNewTaskEngine(t *testing.T) {
	engine, status := NewTaskEngine("testengine", 1024)

	if status != nil {
		engine.Exit()
	}
}
