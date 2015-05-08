package main

import (
	"io"
	"net/http"
)

type Server struct {
	Message string
}

func NewServer(message string) *Server {
	return &Server{message}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, s.Message)
}

func main() {
	http.ListenAndServe(":5000", NewServer("Hello there"))
}
