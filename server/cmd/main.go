package main

import (
	"docker_and_actions_tutorial/server/pkg/base/config"
	log "docker_and_actions_tutorial/server/pkg/base/logging"
	"docker_and_actions_tutorial/server/pkg/server"
)

func init() {
	log.LoggerSetup()
	config.Init()
}

func main() {
	server.Run()
}
