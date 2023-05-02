package pkg

import (
	"math/rand"
	"time"
)

// GetRandomDuration returns a time between 1ms and 2000ms
func GetRandomDuration() time.Duration {
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 2000

	return time.Duration(rand.Intn(max)+min) * time.Millisecond
}
