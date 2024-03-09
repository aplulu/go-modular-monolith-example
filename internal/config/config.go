package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Listen                string `envconfig:"listen" default:""`
	Port                  string `envconfig:"port" default:"8080"`
	GRPCReflectionService bool   `envconfig:"grpc_reflection_service" default:"true"`
}

var conf config

// LoadConf loads the configuration from the environment variables.
func LoadConf() error {
	if err := envconfig.Process("", &conf); err != nil {
		return fmt.Errorf("config.LoadConf: failed to load conf: %w", err)
	}

	return nil
}

func Listen() string {
	return conf.Listen
}

func Port() string {
	return conf.Port
}

func GRPCReflectionService() bool {
	return conf.GRPCReflectionService
}
