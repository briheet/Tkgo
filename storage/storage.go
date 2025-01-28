package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/briheet/tkgo/types"
)

func NewStorage() *NonPresistentMap {
	return &NonPresistentMap{
		Map: make(map[string]UserData),
	}
}

func (s *NonPresistentMap) GetUser(req types.UserTokenRequest) (map[string]any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userData, exists := s.Map[req.UserInfo.UserId]
	if !exists {
		mpp := map[string]any{"token": ""}
		return mpp, false
	}

	minToken := ""
	minValue := int(^uint(0) >> 1)

	for token, count := range userData.TokenCount {
		if count < minValue {
			minValue = count
			minToken = token
		}
	}

	userData.SimulationCount++

	// End of Simulation, need to return all data back
	if userData.SimulationTime >= userData.SimulationCount {
		tokenUsage := []string{}
		leastUsedTokens := []string{}

		for token, count := range userData.TokenCount {
			tokenUsage = append(tokenUsage, fmt.Sprintf("Token %s: %d uses", token, count))
			if count == minValue {
				leastUsedTokens = append(leastUsedTokens, fmt.Sprintf("Token %s (%d use)", token, count))
			}
		}

		return map[string]any{
			"message":         "Simulation Ended",
			"simulationTime":  userData.SimulationTime,
			"tokenUsage":      tokenUsage,
			"leastUsedTokens": leastUsedTokens,
		}, true

	}

	userData.TokenCount[minToken]++

	mpp := map[string]any{"token": minToken}
	return mpp, true
}

func (s *NonPresistentMap) CheckUserPresentOrNot(userId string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Map[userId]

	return exists
}

func (s *NonPresistentMap) CreateANewUser(req types.CreateUser) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tokenCount := make(map[string]int, req.TokenNumbers)
	for i := 1; i <= req.TokenNumbers; i++ {
		tokenName := "Token_" + fmt.Sprint(i)
		tokenCount[tokenName] = 0
	}

	s.Map[req.UserInfo.UserId] = UserData{
		SimulationTime:  req.SimulationTime,
		TimeCreated:     time.Now(),
		TokenCount:      tokenCount,
		SimulationCount: 0,
	}

	// Need to check if the user has been created or not
	exists := s.CheckUserPresentOrNot(req.UserInfo.UserId)
	if !exists {
		return errors.New("not able to create a new User")
	}

	return nil
}
