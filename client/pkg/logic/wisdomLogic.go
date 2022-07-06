package logic

import (
	"net"
	"strconv"

	log "docker_and_actions_tutorial/client/pkg/base/logging"
	"docker_and_actions_tutorial/client/pkg/repository/connectionRepo"
)

// MakeInitialReqForWow requests for service from server
// for this particular project it makes the first contact with server
// and gets the challange
func MakeInitialReqForWow(conn net.Conn) (challange string, err error) {
	// Send request to server
	err = connectionRepo.SendMessage(conn, "")
	if err != nil {
		log.Error().Err(err).Msg("Error on sending initial request")
		return challange, err
	}

	// Read server response for challange
	challange, err = connectionRepo.ReadResponse(conn)
	if err != nil {
		log.Error().Err(err).Msg("Error on reading server response")
		return challange, err
	}

	return challange, err
}

// ProveWork proves that the client is a real client
// by sending the proof to the server
func ProveWork(conn net.Conn, challange string) (err error) {
	// Convert type of challange from string to uint64
	challangeInt, err := strconv.ParseUint(challange, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("Error on parsing challange")
		return err
	}

	// Find all divisors of challange
	divisors := findDPrimeDivisors(challangeInt)
	log.Trace().Interface("Divisors", divisors).Msg("Found divisors")

	// Prepare proof message for server
	proof := prepareProofMessageForServer(divisors)

	// Send divisors to server
	err = connectionRepo.SendMessage(conn, proof)
	if err != nil {
		log.Error().Err(err).Msg("Error on sending divisors")
		return err
	}

	return err
}

// prepareProofMessageForServer prepares the message to let server parse message easier
func prepareProofMessageForServer(divisors []uint64) (proof string) {
	for _, divisor := range divisors {
		proof += strconv.FormatUint(divisor, 10) + " "
	}
	return proof
}

// findDPrimeDivisors finds all divisors of a number
func findDPrimeDivisors(n uint64) (divisors []uint64) {
	var divisor uint64

	for divisor = 2; divisor <= n; divisor++ {
		for n%divisor == 0 {
			divisors = append(divisors, divisor)
			n = n / divisor
		}
	}

	return divisors
}
