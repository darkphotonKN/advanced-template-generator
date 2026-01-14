package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Manager struct {
	projectPath string
}

func NewManager(projectPath string) *Manager {
	return &Manager{projectPath: projectPath}
}

func (m *Manager) Initialize(initialCommitMessage string) error {
	// Change to project directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	if err := os.Chdir(m.projectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}
	defer os.Chdir(originalDir)

	// Remove any existing .git directory
	gitDir := filepath.Join(m.projectPath, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		if err := os.RemoveAll(gitDir); err != nil {
			return fmt.Errorf("failed to remove existing .git directory: %w", err)
		}
	}

	// Initialize git repository
	if err := runCommand("git", "init"); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	// Add all files
	if err := runCommand("git", "add", "."); err != nil {
		return fmt.Errorf("failed to add files to git: %w", err)
	}

	// Create initial commit
	if err := runCommand("git", "commit", "-m", initialCommitMessage); err != nil {
		return fmt.Errorf("failed to create initial commit: %w", err)
	}

	return nil
}

func (m *Manager) IsGitAvailable() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}