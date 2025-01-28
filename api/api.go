package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/briheet/tkgo/storage"
	"github.com/briheet/tkgo/types"
	"go.uber.org/zap"
)

type Server struct {
	ctx     context.Context
	logger  *zap.Logger
	addr    string
	storage *storage.NonPresistentMap
}

func NewServer(ctx context.Context, logger *zap.Logger, addr string, storage *types.NonPresistentMap) *Server {
	return &Server{
		ctx:     ctx,
		logger:  logger,
		addr:    addr,
		storage: storage,
	}
}

func (s *Server) Serve(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", s.Health)
	mux.HandleFunc("GET /getToken", s.GetToken)
}

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
		return
	}
}
