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
	userData, exists := s.storage.GetUser(requestData)
	if !exists {
		s.logger.Error("User not present in memory", zap.Error(err))
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The user not present in memory, please make a new user"))
		return
	}

	// The user is present, return the token, we have already incremented its value
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": userData})
}
