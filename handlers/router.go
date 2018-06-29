package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"gitlab.com/Prophesians/swivel-server/context"
)

func Router(appContext *context.AppContext) http.Handler {
	router := mux.NewRouter()
	swivelServerHandler := SwivelServerHandler{}

	router.HandleFunc("/api", swivelServerHandler.IsAlive()).Methods("GET")
	router.HandleFunc("/api/{user_id}/articles", swivelServerHandler.GetArticles()).Methods("GET")
	router.HandleFunc("/api/tags", swivelServerHandler.SaveTags(appContext)).Methods("POST")

	return router
}