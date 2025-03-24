package cmd

import (
	"fmt"

	"github.com/anthony-magana/git-secrets-scanner/internal/config"

	"github.com/spf13/cobra"
)

// listPatternsCmd represents the list-patterns command
var listPatternsCmd = &cobra.Command{
	Use:   "list-patterns",
	Short: "Lists all regex patterns used for scanning",
	Long: `This command displays the regex patterns currently used by Git Secrets Scanner. 
Patterns are loaded from config.yaml, or defaults are used if no config file is found.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("⚠️ Warning: Using default regex patterns (config.yaml not found or invalid)")
			cfg = config.DefaultConfig()
		}

		// Print loaded patterns
		fmt.Println("🔍 Active Regex Patterns:")
		for _, pattern := range cfg.Patterns {
			fmt.Println("  -", pattern)
		}
	},
}

func init() {
	rootCmd.AddCommand(listPatternsCmd)
	listPatternsCmd.Flags().StringVarP(&configPath, "config", "c", "config.yaml", "Path to custom config file")
}
