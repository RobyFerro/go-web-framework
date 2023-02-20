package interactors

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RobyFerro/go-web-framework/helpers"
)

// GenerateAppKey will generates new application key
type GenerateAppKey struct{}

// Call execute GenerateAppKey interactors
func (c GenerateAppKey) Call() {
	path := helpers.GetDynamicPath("config/server.go")
	read, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	appKey, err := c.generateNewToken()
	if err != nil {
		log.Fatal(err)
	}

	newContent := strings.Replace(string(read), "REPLACE_WITH_CUSTOM_APP_KEY", appKey, -1)
	if err = os.WriteFile(path, []byte(newContent), 0); err != nil {
		log.Fatal(err)
	}
}

func (c GenerateAppKey) generateNewToken() (string, error) {
	data := make([]byte, 10)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("%x", hash[:])

	return hashStr, nil
}
