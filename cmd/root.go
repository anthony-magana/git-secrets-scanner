package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-secrets-scanner",
	Short: "A CLI tool to scan staged Git files for secrets",
	Long: `Git Secrets Scanner is a pre-commit hook and CLI tool that helps 
prevent accidental commits of sensitive credentials (e.g., API keys, tokens).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Git Secrets Scanner! Use --help to see available commands.")
	},
}

// Execute runs the root command and all subcommands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define global flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
}
