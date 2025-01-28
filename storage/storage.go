package storage

import "github.com/briheet/tkgo/types"

func NewStorage() *NonPresistentMap {
	return &NonPresistentMap{
		Map: make(map[string]UserData),
	}
}

func (s *NonPresistentMap) GetUser(types.UserTokenRequest) (string, bool) {
	return "", true
}
