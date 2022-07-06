package logic

import (
	"testing"
	
)

func TestCheckClientProof(t *testing.T) {
	answer := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	proof := "1 2 3 4 5 6 7 8 9 10"
	isCorrect, err := CheckClientProof(answer, proof)
	if err != nil {
		t.Error(err)
	}
	if !isCorrect {
		t.Error("Proof is not correct")
	}
}
