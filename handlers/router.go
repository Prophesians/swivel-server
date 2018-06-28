package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()
	swivelServerHandler := SwivelServerHandler{}

	router.HandleFunc("/", swivelServerHandler.IsAlive()).Methods("GET")
	router.HandleFunc("/topics", swivelServerHandler.GetTopics()).Methods("GET")
	router.HandleFunc("/tags", swivelServerHandler.SaveTags()).Methods("POST")

	return router
}