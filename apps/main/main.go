package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"github.com/Prophesians/swivel-server/handlers"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(handlers.Router())

	http.ListenAndServe(":8000", n)
}
