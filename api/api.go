package api

import (
	"fmt"
	"github.com/skullkon/info-app/pkg/logging"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	logger     *logging.Logger
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		},
	}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
