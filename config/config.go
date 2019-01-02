package config

import (
	"encoding/json"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Configuration structure which hold information for configuring the datasetAPI
type Configuration struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	Host                    string        `envconfig:"HOST_NAME"`
	MongoConfig             *MongoConfig
}

// MongoConfig structure which contains information to access mongo datastore
type MongoConfig struct {
	BindAddr   string `envconfig:"MONGO_ADDR"`
	Collection string `envconfig:"MONGO_COLLECTION"`
	Database   string `envconfig:"MONGO_DATABASE"`
}

var cfg *Configuration

// Get the application and returns the configuration structure
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                ":10000",
		GracefulShutdownTimeout: 5 * time.Second,
		Host:                    "http://localhost",
		MongoConfig: &MongoConfig{
			BindAddr:   "localhost:27017",
			Collection: "courses",
			Database:   "courses",
		},
	}

	return cfg, envconfig.Process("", cfg)
}

// String is implemented to prevent sensitive fields being logged.
// The config is returned as JSON with sensitive fields omitted.
func (config Configuration) String() string {
	json, _ := json.Marshal(config)
	return string(json)
}
