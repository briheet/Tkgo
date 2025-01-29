package storage

import (
	"fmt"
	"time"

	"github.com/briheet/tkgo/types"
)

type Storage interface {
	GetUserToken(types.User) (map[string]any, bool)
	CheckUserPresentOrNot(string) bool
	CreateANewUser(types.CreateUser) error
}

func NewStorage() *NonPresistentMap {
	return &NonPresistentMap{
		Map: make(map[string]UserData),
	}
}

func (s *NonPresistentMap) GetUserToken(req types.User) (map[string]any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userData, exists := s.Map[req.UserId]
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

	totalTokensUsed := 0
	for _, count := range userData.TokenCount {
		totalTokensUsed += count
	}

	if userData.SimulationDone {
		return map[string]any{
			"message":         "No more tokens available, simulation limit reached",
			"simulationTime":  userData.SimulationTime,
			"totalTokensUsed": totalTokensUsed,
		}, true
	}

	// End of Simulation, need to return all data back
	if userData.SimulationTime == totalTokensUsed {
		tokenUsage := []string{}
		leastUsedTokens := []string{}

		for token, count := range userData.TokenCount {
			tokenUsage = append(tokenUsage, fmt.Sprintf("Token %s: %d uses", token, count))
			if count == minValue {
				leastUsedTokens = append(leastUsedTokens, fmt.Sprintf("Token %s (%d use)", token, count))
			}
		}

		userData.SimulationDone = true
		s.Map[req.UserId] = userData

		return map[string]any{
			"message":         "Simulation Ended",
			"simulationTime":  userData.SimulationTime,
			"tokenUsage":      tokenUsage,
			"leastUsedTokens": leastUsedTokens,
		}, true

	}

	// if userData.SimulationCount > userData.SimulationTime {
	// 	mpp := map[string]any{"message": "the SimulationCount has exceeded simulationTime, wait for the tokens to be refreshed"}
	// 	return mpp, false
	// }

	// fmt.Println("Stored users in memory:", s.Map)
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
		UserName:       req.UserInfo.UserName,
		SimulationTime: req.SimulationTime,
		TimeCreated:    time.Now(),
		TokenCount:     tokenCount,
		SimulationDone: false,
	}

	return nil
}
