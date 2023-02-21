package interactors

import (
	"os/exec"
)

// UpdateAlfred updates Alfred cli
type UpdateAlfred struct{}

// Call update Alfred command
func (c UpdateAlfred) Call() error {
	cmd := exec.Command("go", "install", "./cmd/alfred/...")
	return cmd.Run()
}
