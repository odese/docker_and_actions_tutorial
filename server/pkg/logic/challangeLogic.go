package logic

import "crypto/rand"

func PrepareChallange() (challange uint64, answer []uint64) {
	x, _ := rand.Prime(rand.Reader, 30)
	y, _ := rand.Prime(rand.Reader, 30)

	z := x.Uint64() * y.Uint64()

	challange = z
	answer = append(answer, x.Uint64())
	answer = append(answer, y.Uint64())

	return challange, answer
}
