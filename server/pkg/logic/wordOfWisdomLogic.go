package logic

import (
	"docker_and_actions_tutorial/server/pkg/utils"
	math "math/rand"
)

func PickRandomWordOfWisdom() string {
	randomIndex := math.Intn(len(utils.WoW))
	pick := utils.WoW[randomIndex]
	return pick
}
