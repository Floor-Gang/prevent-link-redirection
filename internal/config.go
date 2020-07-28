package internal

import (
	utilConfig "github.com/Floor-Gang/utilpkg/config"
	"log"
)

// Config structure.
type Config struct {
	Auth                string        `yaml:"auth_server"` // authserver ip and port
	Prefix              string        `yaml:"prefix"`      // prefix associated with this bot
	NotificationChannel string        `yaml:"NotificationChannel"`     // channel to report to
}

const configPath = "./config.yml"

// GetConfig retrieves a configuration.
func GetConfig() Config {
	config := Config{
		Auth:                "",
		Prefix:              ".stop_redirect",
		NotificationChannel: "",
	}
	err := utilConfig.GetConfig(configPath, &config)

	if err != nil {
		log.Fatalln(err)
	}

	return config
}

// Save saves configuration
func (config *Config) Save() {
	if err := utilConfig.Save(configPath, config); err != nil {
		log.Println("Failed to save config", err)
	}
}
