package logic

import (
	"testing"
)

func TestPrepareChallange(t *testing.T) {
	challange, answer := PrepareChallange()
	t.Logf("challange: %d, answer: %v", challange, answer)
}
