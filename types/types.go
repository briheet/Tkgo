package types

import (
	"sync"
	"time"
)

type User struct {
	UserName string `json:"username"`
	UserId   string `json:"userid"`
}

type MapData struct {
	UserInfo  User
	TokenData map[string]int
}

type UserTokenRequest struct {
	UserInfo User
	TokenID  string `json:"tokenID"`
}

type UserData struct {
	SimulationTime int
	TimeCreated    time.Time
	TokenCount     map[string]int
}

type NonPresistentMap struct {
	mu  sync.RWMutex
	Map map[string]UserData
}
