package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// installHookCmd represents the install-hook command
var installHookCmd = &cobra.Command{
	Use:   "install-hook",
	Short: "Installs Git pre-commit hook to scan for secrets",
	Long: `This command installs a Git pre-commit hook that runs 
git-secrets-scanner before each commit to prevent accidental secret leaks.`,
	Run: func(cmd *cobra.Command, args []string) {
		installGitHook()
	},
}

func init() {
	rootCmd.AddCommand(installHookCmd)
}

// installGitHook sets up the pre-commit hook
func installGitHook() {
	// Ensure we're inside a Git repository
	if _, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output(); err != nil {
		fmt.Println("Error: Not inside a Git repository.")
		return
	}

	hookPath := filepath.Join(".git", "hooks", "pre-commit")
	hookContent := `#!/bin/sh
# Git Secrets Scanner Pre-Commit Hook
git-secrets-scanner scan
if [ $? -ne 0 ]; then
  echo "❌ Commit blocked! Remove secrets before committing."
  exit 1
fi
`

	// Write the pre-commit hook file
	if err := os.WriteFile(hookPath, []byte(hookContent), 0755); err != nil {
		fmt.Printf("Error installing Git hook: %v\n", err)
		return
	}

	fmt.Println("✅ Git pre-commit hook installed successfully!")
}
