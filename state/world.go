package state

import (
	"time"
)

type World struct {
	UpdateInterval time.Duration
	LastTick       time.Time
}

var WorldInstance = &World{
	UpdateInterval: 300 * time.Millisecond,
	LastTick:       time.Now(),
}
