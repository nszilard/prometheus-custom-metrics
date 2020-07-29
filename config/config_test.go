package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	defer func() {
		os.Setenv("APPLICATION_ENV", "")
	}()

	cases := []struct {
		env string
	}{
		{
			env: Testing,
		},
		{
			env: Production,
		},
	}
	for _, c := range cases {
		os.Setenv("APPLICATION_ENV", c.env)
		actual := GetConfig()
		assert.Equal(t, actual.Environment, c.env)
		conf.IsSetup = false
	}
}
