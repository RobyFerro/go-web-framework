package services

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

// GitServicesImpl wraps git library
type GitServicesImpl struct{}

// Clone specific repository
func (c GitServicesImpl) Clone(repo, destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      repo,
		Progress: nil,
	})

	return err
}

// Remove git directory from specic repo
func (c GitServicesImpl) Remove(repo string) error {
	path := fmt.Sprintf("%s/.git", repo)
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = repo
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// Update specific repository
func (c GitServicesImpl) Update(repo, destination string) error {
	cmd := exec.Command("go", "get", "-u", repo)
	cmd.Dir = destination

	return cmd.Run()
}
