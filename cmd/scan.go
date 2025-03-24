package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/anthony-magana/git-secrets-scanner/internal/config"
	"github.com/anthony-magana/git-secrets-scanner/internal/git"
	"github.com/anthony-magana/git-secrets-scanner/internal/scanner"

	"github.com/spf13/cobra"
)

var (
	excludePatterns []string
	configPath      string
	verbose         bool
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans staged Git files for secrets",
	Long: `This command scans the currently staged files in a Git repository
to detect secrets such as API keys, passwords, and tokens before they are committed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load config
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("⚠️ Warning: Using default settings (config.yaml not found or invalid)")
			cfg = config.DefaultConfig()
		}

		// Get staged files
		stagedFiles, err := git.GetStagedFiles()
		if err != nil {
			log.Fatalf("Error getting staged files: %v", err)
		}

		if len(stagedFiles) == 0 {
			fmt.Println("No staged files found.")
			return
		}

		fmt.Println("Scanning staged files for secrets...")

		// Scan each file
		for _, file := range stagedFiles {
			if isExcluded(file) {
				if verbose {
					fmt.Printf("ℹ️ Skipping excluded file: %s\n", file)
				}
				continue
			}

			fmt.Printf("🔍 Scanning: %s\n", file)
			foundSecrets := scanner.ScanFileWithEntropy(file, cfg.EntropyThreshold, cfg.Patterns)

			if len(foundSecrets) > 0 {
				fmt.Printf("❌ Potential secrets found in %s:\n", file)
				for _, secret := range foundSecrets {
					fmt.Println("  -", secret)
				}
			} else {
				if verbose {
					fmt.Println("✅ No secrets detected in", file)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringSliceVarP(&excludePatterns, "exclude", "e", []string{}, "Files or patterns to exclude (e.g., config.json, *.log)")
	scanCmd.Flags().StringVarP(&configPath, "config", "c", "config.yaml", "Path to custom config file")
	scanCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

// isExcluded checks if a file matches any exclusion pattern
func isExcluded(filePath string) bool {
	for _, pattern := range excludePatterns {
		matched, _ := filepath.Match(pattern, filepath.Base(filePath))
		if matched || strings.HasSuffix(filePath, pattern) {
			return true
		}
	}
	return false
}
