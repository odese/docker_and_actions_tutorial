package logic

import (
	"strconv"
	"strings"

	log "docker_and_actions_tutorial/server/pkg/base/logging"
	"docker_and_actions_tutorial/server/pkg/utils"
)

// CheckClientProof checks the client proof
func CheckClientProof(answer []uint64, proofMessage string) (isCorrect bool, err error) {
	proof := make([]uint64, 0)
	clientProof := strings.Split(proofMessage, " ")
	for _, v := range clientProof {
		p, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("Error on parsing client proof.")
			return isCorrect, err
		}
		proof = append(proof, p)
	}

	if utils.CompareTwoArrays(answer, proof) {
		isCorrect = true
	}

	return isCorrect, err
}
