package tool

import (
	"os"
	"path/filepath"
)

// GetDynamicPath returns the absolute path of the selected file/folder.
// The basic path is Go-Web main folder.
// Example: GetDynamicPath("storage/certs/tls.key")
func GetDynamicPath(path string) string {
	test := os.Getenv("base_path")
	return filepath.Join(test, path)
}
