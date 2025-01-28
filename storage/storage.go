package storage

import (
	"github.com/briheet/tkgo/types"
)

func NewStorage() *NonPresistentMap {
	return &NonPresistentMap{
		Map: make(map[string]UserData),
	}
}

func (s *NonPresistentMap) GetUser(req types.UserTokenRequest) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userData, exists := s.Map[req.UserInfo.UserId]
	if !exists {
		return "", false
	}

	minToken := ""
	minValue := int(^uint(0) >> 1)

	for token, count := range userData.TokenCount {
		if count < minValue {
			minValue = count
			minToken = token
		}
	}

	if minToken == "" {
		return "", false
	}

	userData.TokenCount[minToken]++

	return minToken, true
}
