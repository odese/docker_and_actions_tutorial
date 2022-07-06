package controller

import (
	"net"
	"strconv"

	log "docker_and_actions_tutorial/server/pkg/base/logging"
	"docker_and_actions_tutorial/server/pkg/logic"
	"docker_and_actions_tutorial/server/pkg/repository/connectionRepo"
)

func WordOfWisdomController(conn net.Conn, req string) {
	challenge, answer := logic.PrepareChallange()

	// Send challange to client
	err := connectionRepo.SendMessage(conn, strconv.FormatUint(challenge, 10))
	if err != nil {
		log.Error().Err(err).Msg("Error on sending challenge to client.")
		return
	} else {
		log.Trace().
			Interface("To", conn.RemoteAddr()).
			Uint64("Challenge", challenge).
			Msg("Challange sent to client.")
	}

	// Read client proof
	proofMessage, err := connectionRepo.ReadResponse(conn)
	if err != nil {
		log.Error().Err(err).Msg("Error on reading proof message from client.")
		return
	}

	// Check client proof
	isCorrect, err := logic.CheckClientProof(answer, proofMessage)
	if err != nil {
		log.Error().Err(err).Msg("Error on checking client proof.")
		return
	}

	// Send response to client
	if isCorrect {
		log.Trace().
			Interface("Client", conn.RemoteAddr()).
			Msg("Client proof is correct.")
		connectionRepo.SendMessage(conn, logic.PickRandomWordOfWisdom())
	} else {
		log.Trace().
			Interface("Client", conn.RemoteAddr()).
			Msg("Client proof is incorrect.")
		connectionRepo.SendMessage(conn, "FAIL")
	}
}
