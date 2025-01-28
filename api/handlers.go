package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/briheet/tkgo/types"
	"go.uber.org/zap"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Running health handler")
	fmt.Fprintf(w, "Running Good")
}

func (s *Server) GetToken(w http.ResponseWriter, r *http.Request) {
	var requestData types.UserTokenRequest

	// Check body data issues
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		s.logger.Error("Failed to decode request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invaid data format"))
		return
	}

	// Check if current inmemory map has the user map or not
	// if yes, get the data back
	// the data can be 2 types
	// either the token which is great, else all the token count as the simulation has ended
	userData, exists := s.storage.GetUser(requestData)
	if !exists {
		s.logger.Error("User not present in memory", zap.Error(err))
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The user not present in memory, please make a new user"))
		return
	}

	// The user is present, return the token, we have already incremented its value
	s.logger.Info("The token has been accessed successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
	// json.NewEncoder(w).Encode(map[string]string{"token": userData})
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestUserData types.CreateUser

	err := json.NewDecoder(r.Body).Decode(&requestUserData)
	if err != nil {
		s.logger.Error("Failed to decode request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invaid data format"))
		return
	}

	// Checks if the user has already tokens present or not
	// if yes, return as we do not need to create a new user
	// else procees
	exists := s.storage.CheckUserPresentOrNot(requestUserData.UserInfo.UserId)
	if exists {
		s.logger.Error("The user is already present")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("The user already exists in memory, cannot create a new user"))
		return
	}

	err = s.storage.CreateANewUser(requestUserData)
	if err != nil {
		s.logger.Error("Failed to create a new user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("The user has not been created due to internal server errors"))
		return
	}

	s.logger.Info("A user has been successfully created")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
