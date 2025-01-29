package storage

import (
	"sync"
	"time"
)

type NonPresistentMap struct {
	Mu  sync.RWMutex
	Map map[string]UserData
}

type UserData struct {
	UserName       string
	SimulationTime int
	TimeCreated    time.Time
	TokenCount     map[string]int
	SimulationDone bool
}
