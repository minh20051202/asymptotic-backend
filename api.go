package main

import "net/http"

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
