package storage

import (
	"testing"

	"github.com/briheet/tkgo/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateANewUser(t *testing.T) {
	storage := NewStorage()
	req := types.CreateUser{
		UserInfo: types.User{
			UserName: "TestUser",
			UserId:   "user1",
		},
		TokenNumbers:   3,
		SimulationTime: 10,
	}

	err := storage.CreateANewUser(req)
	assert.Nil(t, err, "Expected no error when creating a new user")

	assert.True(t, storage.CheckUserPresentOrNot("user1"), "User should be present in storage")
}

func TestCheckUserPresentOrNot(t *testing.T) {
	storage := NewStorage()
	assert.False(t, storage.CheckUserPresentOrNot("user1"), "User should not be present initially")

	req := types.CreateUser{
		UserInfo: types.User{
			UserName: "TestUser",
			UserId:   "user1",
		},
		TokenNumbers:   3,
		SimulationTime: 10,
	}
	_ = storage.CreateANewUser(req)

	assert.True(t, storage.CheckUserPresentOrNot("user1"), "User should be present after creation")
}

func TestGetUserToken(t *testing.T) {
	storage := NewStorage()

	req := types.CreateUser{
		UserInfo: types.User{
			UserId:   "user1",
			UserName: "TestUser",
		},
		TokenNumbers:   2,
		SimulationTime: 5,
	}

	_ = storage.CreateANewUser(req)

	tokenResponse, success := storage.GetUserToken(types.User{UserId: "user1"})
	assert.True(t, success, "Expected to get a token for existing user")
	assert.NotEqual(t, "", tokenResponse["token"], "Token should not be empty")

	// chekc non-existing user
	_, success = storage.GetUserToken(types.User{UserId: "user2"})
	assert.False(t, success, "Should return false for a non-existing user")
}

func TestSimulationEndCondition(t *testing.T) {
	storage := NewStorage()

	req := types.CreateUser{
		UserInfo: types.User{
			UserId:   "user1",
			UserName: "TestUser",
		},
		TokenNumbers:   1,
		SimulationTime: 2,
	}

	_ = storage.CreateANewUser(req)

	// Use both available tokens
	storage.GetUserToken(types.User{UserId: "user1"})
	storage.GetUserToken(types.User{UserId: "user1"})

	// Simulation ends by now
	tokenResponse, success := storage.GetUserToken(types.User{UserId: "user1"})

	assert.True(t, success, "Expected success after simulation end")
	assert.Equal(t, "Simulation Ended", tokenResponse["message"], "Should indicate simulation has ended")
}
