package server

import (
	"bufio"
	"io"
	"net"

	"docker_and_actions_tutorial/server/pkg/base/config"
	log "docker_and_actions_tutorial/server/pkg/base/logging"
	"docker_and_actions_tutorial/server/pkg/controller"
)

func Run() {
	network := config.GetConfig().GetString("server.network")
	address := config.GetConfig().GetString("server.address")
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal().Err(err).Msg("Error on starting server.")
	} else {
		log.Info().Msg("Listening and serving TCP on " + listener.Addr().String())
	}
	defer listener.Close()

	// Waiting for a client to connect
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error().Err(err).Msg("Error on accepting connection")
			continue
		} else {
			log.Info().Interface("Client", conn.RemoteAddr()).Msg("Accepted connection from new client.")
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientReader := bufio.NewReader(conn)

	// Waiting for the client request
	for {
		clientRequest, err := clientReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Trace().
					Interface("Client", conn.RemoteAddr()).
					Msg("Client terminated the connection")
				return
			} else {
				log.Error().
					Interface("Client", conn.RemoteAddr()).
					Err(err).
					Msg("Error on reading client request")
				return
			}
		}

		controller.WordOfWisdomController(conn, clientRequest)
	}
}
