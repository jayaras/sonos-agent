package config

import (
	"os"
	"strings"
)

type Config struct {
}

func (c *Config) GetString(name string, def string) string {

	v, b := os.LookupEnv(strings.ToUpper(name))

	if b {
		return v
	} else {
		return def
	}

}

func (c *Config) GetTCPConnection(name string, def string) string {
	h := c.GetString(name, def)
	u := "tcp://" + h
	return u
}
