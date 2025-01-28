package storage

import (
	"sync"
	"time"
)

type NonPresistentMap struct {
	mu  sync.RWMutex
	Map map[string]UserData
}

type UserData struct {
	SimulationTime  int
	TimeCreated     time.Time
	TokenCount      map[string]int
	SimulationCount int
}
