package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"gitlab.com/Prophesians/swivel-server/handlers"
	"gitlab.com/Prophesians/swivel-server/config"
	"flag"
	"log"
	"strconv"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(handlers.Router())
	configPath := flag.String("configPath", "./config/config.development.json", "provide the config path")
	if *configPath == "" {
		log.Fatal("Config file path should be provided")
	}
	appConfig, err := config.GetConfig(*configPath)
	if err != nil {
		log.Fatal("error generating config")
	}
	http.ListenAndServe(":"+strconv.Itoa(appConfig.GetPort()), n)
}
