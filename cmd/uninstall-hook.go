package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// uninstallHookCmd represents the uninstall-hook command
var uninstallHookCmd = &cobra.Command{
	Use:   "uninstall-hook",
	Short: "Removes the Git pre-commit hook",
	Long:  `This command removes the Git pre-commit hook installed by git-secrets-scanner.`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstallGitHook()
	},
}

func init() {
	rootCmd.AddCommand(uninstallHookCmd)
}

// uninstallGitHook removes the pre-commit hook
func uninstallGitHook() {
	hookPath := filepath.Join(".git", "hooks", "pre-commit")

	// Check if the file exists
	if _, err := os.Stat(hookPath); os.IsNotExist(err) {
		fmt.Println("❌ No Git pre-commit hook found to uninstall.")
		return
	}

	// Remove the hook
	if err := os.Remove(hookPath); err != nil {
		fmt.Printf("Error removing Git hook: %v\n", err)
		return
	}

	fmt.Println("✅ Git pre-commit hook uninstalled successfully!")
}
