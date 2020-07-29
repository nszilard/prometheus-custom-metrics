package random

import (
	"math/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 0; i < 6; i++ {
		res, err := Generate()
		if res == nil && err == nil {
			t.Errorf("Generate() => didn't return any value nor any errors")
		}
	}
}

func Test_randomAlphaNumeric(t *testing.T) {
	actual := randomAlphaNumeric(rand.New(rand.NewSource(1)))
	if len(actual) != 34 {
		t.Errorf("randomAlphaNumeric() => Didn't return expected length")
	}
}
