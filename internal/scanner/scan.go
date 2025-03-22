package scanner

import (
	"bufio"
	"os"
	"regexp"
)

// List of regex patterns for secrets
var patterns = []string{
	`AKIA[0-9A-Z]{16}`, // AWS Access Key
	`xox[baprs]-[0-9]{12}-[0-9]{12}-[0-9A-Za-z]{24}`,                  // Slack Token
	`(?i)api[-_]?key['"]?\s*[:=]\s*['"]?([A-Za-z0-9_\-]{20,50})['"]?`, // Generic API Key
	`(?i)password\s*=\s*['"]\w+`,                                      // Generic passwords
	`-----BEGIN RSA PRIVATE KEY-----`,                                 // RSA Private keys
}

// ScanFile scans a given file for secrets
func ScanFile(filePath string) []string {
	var foundSecrets []string

	file, err := os.Open(filePath)
	if err != nil {
		return foundSecrets
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, pattern := range patterns {
			re := regexp.MustCompile(pattern)
			if re.MatchString(line) {
				foundSecrets = append(foundSecrets, line)
			}
		}
	}

	return foundSecrets
}
