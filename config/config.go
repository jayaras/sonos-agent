package config

import (
	"os"
	"strconv"
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

func (c *Config) GetInt(name string, def int) int {
	v, b := os.LookupEnv(strings.ToUpper(name))

	if b {
		x, err := strconv.Atoi(v)
		if err == nil {
			return x
		}
	}

	return def

}
