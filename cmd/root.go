package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "git-secrets-scanner",
	Short: "A CLI tool to scan staged Git files for secrets",
	Long: `Git Secrets Scanner detects and prevents accidental commits of sensitive 
credentials (API keys, passwords, tokens) by scanning staged files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'git-secrets-scanner --help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add a custom help flag (optional, since Cobra provides one by default)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help for the CLI tool")
}
