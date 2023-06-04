package http

import (
	"bxb_home/internal/controllers/hello"
	"bxb_home/internal/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
	handler http.Handler
	server  *http.Server
}

func New() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", hello.Handler)
	handler := middleware.Logging(mux)
	server := &http.Server{Addr: ":3000", Handler: handler}

	return &Server{handler: handler, server: server}
}

func (h Server) Start() {
	addr := os.Getenv("ADDR")
	log.Printf("handler is listening at %s", addr)
	log.Fatal(h.server.ListenAndServe())
}

func (h *Server) Stop(ctx context.Context) error {
	log.Println("stopping server...")
	err := h.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("handler shutdown: %w", err)
	}
	return nil
}
