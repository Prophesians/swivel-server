package handlers

import (
	"net/http"
	"io"
	"fmt"
)

type SwivelServerHandler struct {
}

func (ssh *SwivelServerHandler) IsAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	}
}

func (ssh *SwivelServerHandler) GetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This are todays news")
	}
}

func (ssh *SwivelServerHandler) SaveTags() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Data is saved")
	}
}
