package api

import (
	"context"
	"net/http"

	"github.com/briheet/tkgo/storage"
	"go.uber.org/zap"
)

type Server struct {
	ctx     context.Context
	logger  *zap.Logger
	addr    string
	storage *storage.NonPresistentMap
}

func NewServer(ctx context.Context, logger *zap.Logger, addr string, storage *storage.NonPresistentMap) *Server {
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
	mux.HandleFunc("GET /createUser", s.CreateUser)
}
