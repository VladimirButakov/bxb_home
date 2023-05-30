package http

import (
	"bxb_home/internal/controllers/hello"
	"bxb_home/internal/middleware"
	"log"
	"net/http"
	"os"
)

type Server struct {
	server http.Handler
}

func New() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", hello.Handler)
	handler := middleware.Logging(mux)

	return &Server{server: handler}
}

func (h Server) Start() {
	addr := os.Getenv("ADDR")
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(":3000", h.server))
}
