# Git Secrets Scanner
### (**STILL IN DEVELOPMENT**)
**Git Secrets Scanner** is a CLI tool designed to detect and prevent accidental commits of sensitive information such as API keys, passwords, and tokens. It integrates with Git hooks to block risky commits before they reach your repository.

## 🚀 Features
- 📜 **Scans staged files** for secrets before committing.
- 🔍 **Regex-based detection** for common API keys and tokens (AWS, Slack, etc.).
- 🧮 **Entropy analysis** to reduce false positives.
- 🛑 **Git Hook integration** to block commits containing secrets.
- 🛠 **Customizable regex patterns** via a configuration file.

---

## 📦 Installation
You can install Git Secrets Scanner by building from source:

```sh
git clone https://github.com/anthony-magana/git-secrets-scanner.git
cd git-secrets-scanner
go build -o git-secrets-scanner
mv git-secrets-scanner /usr/local/bin/
```

#### Alternatively, install **via go install**:

```sh
go install github.com/anthony-magana/git-secrets-scanner@latest
```

## 🛠 Usage

### Scan for secrets manually

Run the following command to scan your staged files:

```sh
git-secrets-scanner scan
```
If any secrets are detected, they will be displayed in the output.

## 🛑 Git Hook Integration
To prevent committing secrets, install the **pre-commit hook**:

```sh
git-secrets-scanner install-hook
```
This will create a .git/hooks/pre-commit script that automatically runs the scanner before each commit.

### Manually Adding a Pre-Commit Hook
If you’d like to manually add a Git pre-commit hook:

```sh
echo '#!/bin/sh
git-secrets-scanner scan || exit 1' > .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```
Now, any attempt to commit a file containing secrets will be blocked.

### Uninstalling the **pre-commit hook**:
```sh
git-secrets-scanner uninstall-hook
```

## ⚙️ Custom Configuration
By default, **Git Secrets Scanner** loads patterns and entropy settings from ``config.yaml``.
You can specify a custom configuration file using the ``--config`` flag.

### Usage
```sh
git-secrets-scanner scan --config custom-config.yaml
```

### Example ``custom-config.yaml``
```yaml
patterns:
  - "sk_live_[0-9a-zA-Z]{24}"  # Stripe Live API Key
  - "ghp_[0-9a-zA-Z]{36}"  # GitHub Personal Access Token
entropy_threshold: 4.0  # Lower threshold for stricter entropy detection
```

## 📖 Example Output

```
$ git-secrets-scanner scan --config custom-config.yaml
Scanning staged files for secrets...
Scanning: credentials.json
Potential secrets found in credentials.json:
  - [Regex] "sk_live_abc123xyz789..."
  - [Entropy] hjK8@#Dfh99s2!

Commit blocked! Remove sensitive data before committing.
```

## 🔍 Entropy-Based Scanning
Git Secrets Scanner now includes entropy-based scanning to detect highly random strings, which are often API keys, passwords, or tokens.

### How It Works

- Uses Shannon entropy to measure randomness.
- Flags strings with high entropy (default threshold: 4.5).
- Helps detect secrets without predefined regex patterns.
```
$ git-secrets-scanner scan
Scanning staged files for secrets...
Scanning: .env
Potential secrets found in .env:
  - [Regex] AWS_SECRET_KEY="AKIAEXAMPLE1234567890"
  - [Entropy] hjK8@#Dfh99s2!
```

## 🛑 Excluding Files from Scanning
To ignore certain files or patterns, use the --exclude flag:

### Usage
```sh
git-secrets-scanner scan --exclude config.json --exclude "*.log"
```

### Example Output
```
$ git-secrets-scanner scan --exclude secrets.txt --exclude "*.log"
Skipping excluded file: secrets.txt
Skipping excluded file: debug.log
Scanning: app.js
No secrets detected in app.js
```

## 📖 Getting Help
You can view all available commands and options by using the --help flag:

### Display general help
```sh
git-secrets-scanner --help
```

### Display help for a specific command
```sh
git-secrets-scanner scan --help
```

### Example Output
```
$ git-secrets-scanner --help

A CLI tool to scan staged Git files for secrets.

Usage:
  git-secrets-scanner [command]

Available Commands:
  scan           Scans staged Git files for secrets
  install-hook   Installs Git pre-commit hook
  uninstall-hook Removes the Git pre-commit hook
  help           Show help for a command

Flags:
  -h, --help     Show help for the CLI tool
  -v, --verbose  Enable verbose output

Use "git-secrets-scanner [command] --help" for more details about a command.
```

## 🤝 Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue.

## 📜 License
This project is licensed under the MIT License.

## 👨‍💻 Author
Developed by Anthony Magana
GitHub: anthony-magana

