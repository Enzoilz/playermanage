package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Bievenue sur Player Manager API!")
}

func (m *PlayerManager) PlayerHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m.Players)
}