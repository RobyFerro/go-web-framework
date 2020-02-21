package command

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"

	gwf "github.com/RobyFerro/go-web-framework"
)

// GenerateKey will generate Go-Web application key in main config.yml file
type GenerateKey struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *GenerateKey) Register() {
	c.Signature = "generate:key"             // Change command signature
	c.Description = "Genate application key" // Change command description
}

// Run this command
func (c *GenerateKey) Run() {
	path := gwf.GetDynamicPath("config.yml")
	read, err := ioutil.ReadFile(path)

	if err != nil {
		gwf.ProcessError(err)
	}

	appKey, err := generateNewToken()
	if err != nil {
		gwf.ProcessError(err)
	}

	newContent := strings.Replace(string(read), "$$APP_KEY$$", appKey, -1)

	if err = ioutil.WriteFile(path, []byte(newContent), 0); err != nil {
		gwf.ProcessError(err)
	}

	return
}

// generateNewToken will return a random sha256 hash
func generateNewToken() (string, error) {
	data := make([]byte, 10)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("%x", hash[:])

	return hashStr, nil
}
