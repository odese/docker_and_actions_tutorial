package client

import (
	"net"

	"docker_and_actions_tutorial/client/pkg/base/config"
	log "docker_and_actions_tutorial/client/pkg/base/logging"
	"docker_and_actions_tutorial/client/pkg/controller"
)

func Run() {
	network := config.GetConfig().GetString("server.network")
	address := config.GetConfig().GetString("server.address")
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal().Err(err).Msg("Error on connecting server.")
	} else {
		log.Info().Msg("Connection established to server on 6090")
	}
	defer conn.Close()

	// Client sends new request after previous request is responded.
	for {
		response, err := controller.HandleConnectionForWordOfWisdom(conn)
		if err != nil {
			log.Error().Err(err).Msg("Error on handling connection")
		} else {
			log.Trace().Msg("Got Word of Wisdom:" + response)
		}
	}
}
