package config

import "github.com/tkanos/gonfig"

type AppConfig struct {
	Port int
}

func GetConfig(filePath string) (*AppConfig, error) {
	appConfig := AppConfig{}
	err := gonfig.GetConf(filePath, &appConfig)
	return &appConfig, err
}

func (a *AppConfig) GetPort() (int) {
	return a.Port
}
