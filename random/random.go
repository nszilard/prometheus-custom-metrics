package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate will return either a string or an error randomly
func Generate() ([]byte, error) {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Simulate application runtime
	t := seededRand.Intn(3000)
	time.Sleep(time.Duration(t) * time.Millisecond)

	// Randomly return an error
	if seededRand.Intn(2) < 1 {
		return nil, fmt.Errorf("random error occurred")
	}

	return []byte(randomAlphaNumeric(seededRand)), nil
}

func randomAlphaNumeric(r *rand.Rand) string {
	const charset = lowercase + uppercase + numeric

	b := make([]byte, 24)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return fmt.Sprintf("Response: %s", string(b))
}
