package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"gitlab.com/Prophesians/swivel-server/handlers"
	"gitlab.com/Prophesians/swivel-server/config"
	"flag"
	"log"
	"strconv"
	"gitlab.com/Prophesians/swivel-server/repositories"
	"gitlab.com/Prophesians/swivel-server/context"
)

func main() {
	configPath := flag.String("configPath", "./config/config.development.json", "provide the config path")
	if *configPath == "" {
		log.Fatal("Config file path should be provided")
	}
	appConfig, err := config.GetConfig(*configPath)
	if err != nil {
		log.Fatal("error generating config")
	}
	repository, err := repositories.GetRepository(appConfig)
	if err != nil {
		log.Fatal("Error :", err)
	}
	err = repository.Ping()
	if err != nil {
		log.Fatal("Error :", err)
	}

	appContext := &context.AppContext{
		AppConfig:  appConfig,
		Repository: &repository,
	}

	n := negroni.Classic()
	n.UseHandler(handlers.Router(appContext))
	http.ListenAndServe(":"+strconv.Itoa(appConfig.GetPort()), n)
}
