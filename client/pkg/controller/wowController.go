package controller

import (
	"net"

	log "docker_and_actions_tutorial/client/pkg/base/logging"
	"docker_and_actions_tutorial/client/pkg/logic"
	"docker_and_actions_tutorial/client/pkg/repository/connectionRepo"
)

func HandleConnectionForWordOfWisdom(conn net.Conn) (response string, err error) {
	log.Info().Msg("New request is sent.")
	// Send initial request to server
	challange, err := logic.MakeInitialReqForWow(conn)
	if err != nil {
		log.Error().Err(err).Msg("Error on requesting service")
		return response, err
	} else {
		log.Trace().Msg("Got Challange:" + challange)
	}

	// Prove that the client is a real client
	err = logic.ProveWork(conn, challange)
	if err != nil {
		log.Error().Err(err).Msg("Error on proving work")
		return response, err
	}

	// Read server response for word of wisdom
	response, err = connectionRepo.ReadResponse(conn)
	if err != nil {
		log.Error().Err(err).Msg("Error on reading server response")
		return response, err
	}

	return response, err
}
