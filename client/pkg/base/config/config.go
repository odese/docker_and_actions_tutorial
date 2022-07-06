package config

import (
	log "docker_and_actions_tutorial/client/pkg/base/logging"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that starts the viper (external lib)
// and returns the configuration struct.
func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("yml")
	config.SetConfigName(GetConfigName())
	config.AddConfigPath("configs")

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("Error on reading config")
	}
}

//GetConfig ...
func GetConfig() *viper.Viper {
	return config
}

//GetConfigName ...
func GetConfigName() string {
	return "config"
}
