package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config structure for regex patterns
type Config struct {
	Patterns []string `yaml:"patterns"`
}

// LoadConfig reads the config.yaml file
func LoadConfig(configPath string) (Config, error) {
	var cfg Config

	// Open file
	file, err := os.Open(configPath)
	if err != nil {
		return cfg, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	// Decode YAML
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config file: %v", err)
	}

	return cfg, nil
}

// DefaultConfig returns default regex patterns
func DefaultConfig() Config {
	return Config{
		Patterns: []string{
			"AKIA[0-9A-Z]{16}", // AWS Access Key
			"xox[baprs]-[0-9]{12}-[0-9]{12}-[0-9A-Za-z]{24}",                        // Slack Token
			"(?i)api[-_]?key['\"]?\\s*[:=]\\s*['\"]?([A-Za-z0-9_\\-]{20,50})['\"]?", // Generic API Key
		},
	}
}
