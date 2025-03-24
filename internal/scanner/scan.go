package scanner

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// ScanFileWithEntropy scans a file for secrets using regex and entropy checks
func ScanFileWithEntropy(filePath string, entropyThreshold float64, patterns []string) []string {
	var foundSecrets []string

	file, err := os.Open(filePath)
	if err != nil {
		return foundSecrets
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Regex-based scanning
		for _, pattern := range patterns {
			re := regexp.MustCompile(pattern)
			if re.MatchString(line) {
				foundSecrets = append(foundSecrets, "[Regex] "+line)
			}
		}

		// Entropy-based scanning
		for _, word := range splitWords(line) {
			if isHighEntropy(word, entropyThreshold) {
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
