package scanner

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"git-secrets-scanner/internal/config"
)

// ScanFile scans a given file for secrets using regex patterns from config.yaml
func ScanFile(filePath string) []string {
	var foundSecrets []string

	// Load regex patterns from config.yaml (fallback to defaults)
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("⚠️ Warning: Using default regex patterns (config.yaml not found or invalid)")
		cfg = config.DefaultConfig()
	}

	file, err := os.Open(filePath)
	if err != nil {
		return foundSecrets
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, pattern := range cfg.Patterns {
			re := regexp.MustCompile(pattern)
			if re.MatchString(line) {
				foundSecrets = append(foundSecrets, line)
			}
		}
	}

	return foundSecrets
}
