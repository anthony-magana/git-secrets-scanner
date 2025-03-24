package scanner

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"git-secrets-scanner/internal/config"
)

// ScanFile scans a file for secrets using regex and entropy checks
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

		// Regex-based scanning
		for _, pattern := range cfg.Patterns {
			re := regexp.MustCompile(pattern)
			if re.MatchString(line) {
				foundSecrets = append(foundSecrets, "[Regex] "+line)
			}
		}

		// Entropy-based scanning
		for _, word := range splitWords(line) {
			if isHighEntropy(word) {
				foundSecrets = append(foundSecrets, "[Entropy] "+word)
			}
		}
	}

	return foundSecrets
}

// splitWords extracts potential secrets from a line
func splitWords(line string) []string {
	delimiters := " \t=:\"',;(){}[]<>"
	return strings.FieldsFunc(line, func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	})
}
