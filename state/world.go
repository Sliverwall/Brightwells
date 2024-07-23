package state

import (
	"Brightwells/config"
	"time"
)

type World struct {
	UpdateInterval time.Duration
	LastTick       time.Time
}

var WorldInstance = &World{
	UpdateInterval: config.TICK_RATE * time.Millisecond,
	LastTick:       time.Now(),
}
