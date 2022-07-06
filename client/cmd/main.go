package main

import (
	"docker_and_actions_tutorial/client/pkg/base/config"
	log "docker_and_actions_tutorial/client/pkg/base/logging"
	"docker_and_actions_tutorial/client/pkg/client"
)

func init() {
	log.LoggerSetup()
	config.Init()
}

func main() {
	client.Run()
}
