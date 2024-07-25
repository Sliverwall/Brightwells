package systems

import "time"

type TickManager struct {
	LastTick       time.Time
	UpdateInterval time.Duration
}

func NewTickManager(updateInterval time.Duration) *TickManager {
	return &TickManager{
		LastTick:       time.Now(),
		UpdateInterval: updateInterval,
	}
}

func (tm *TickManager) ShouldUpdate() bool {
	currentTime := time.Now()
	if currentTime.Sub(tm.LastTick) >= tm.UpdateInterval {
		tm.LastTick = currentTime
		return true
	}
	return false
}
