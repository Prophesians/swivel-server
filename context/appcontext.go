package context

import (
	"gitlab.com/Prophesians/swivel-server/config"
	"gitlab.com/Prophesians/swivel-server/repositories"
)

type AppContext struct {
	AppConfig *config.AppConfig
	Repository *repositories.Repository
}
