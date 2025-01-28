package api

import (
	"encoding/json"
	"net/http"

	"github.com/briheet/tkgo/types"
	"go.uber.org/zap"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Running health handler")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "server healthy"})
}

func (s *Server) GetToken(w http.ResponseWriter, r *http.Request) {
	var requestData types.User

	// Check body data issues
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		s.logger.Error("Failed to decode request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid data format"})
		return
	}

	// Check if current inmemory map has the user map or not
	// if yes, get the data back
	// the data can be 2 types
	// either the token which is great, else all the token count as the simulation has ended
	userData, exists := s.storage.GetUserToken(requestData)
	if !exists {
		s.logger.Error("User not present in memory", zap.Error(err))
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "The user not present in memory, please make a new user"})
		return
	}

	// The user is present, return the token, we have already incremented its value
	s.logger.Info("The token has been accessed successfully")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
	// json.NewEncoder(w).Encode(map[string]string{"token": userData})
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestUserData types.CreateUser

	err := json.NewDecoder(r.Body).Decode(&requestUserData)
	if err != nil {
		s.logger.Error("Failed to decode request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid data format"})
		return
	}

	// Checks if the user has already tokens present or not
	// if yes, return as we do not need to create a new user
	// else procees
	exists := s.storage.CheckUserPresentOrNot(requestUserData.UserInfo.UserId)
	if exists {
		s.logger.Error("The user is already present")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "The user already exists in memory, cannot create a new user"})
		return
	}

	// Create a new User, and also check if it was created or not
	// We are checking it in the function itself so that there is not again locking of resources
	// that we get after mu.Lock()

	err = s.storage.CreateANewUser(requestUserData)
	if err != nil {
		s.logger.Error("Failed to create a new user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "The user has not been created due to internal server errors"})
		return
	}

	exists = s.storage.CheckUserPresentOrNot(requestUserData.UserInfo.UserId)
	if !exists {
		s.logger.Error("Failed to create a new user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "The user has not been created due to internal server errors"})
		return
	}

	s.logger.Info("A user has been successfully created")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
