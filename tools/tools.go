package tools

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Default struct {
	TimeoutUpdate int  // default: 60 sec
	Debug         bool // default: true
	Token         string
}

// Contains config data
type Config struct {
	Default Default
}

// Return config data
func ParseConfigToml(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var conf Config

	if _, err := toml.Decode(string(data), &conf); err != nil {
		return &Config{}, err
	}

	return &conf, nil
}
