package services

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/helpers"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CommandServiceImpl implement CommandService interface and handles method to creates new cli commands
type CommandServiceImpl struct {
	command entities.Command
}

// SnakeToCamelCase converts snake case string to camel case
func (c CommandServiceImpl) SnakeToCamelCase(raw string) string {
	data := strings.Split(strings.ToLower(raw), "_")
	for i, name := range data {
		data[i] = cases.Title(language.Und, cases.NoLower).String(name)
	}

	return strings.Join(data, "")
}

// ReadBlueprint reads cli blueprints
func (c CommandServiceImpl) ReadBlueprint(name string) string {
	var _, filename, _, _ = runtime.Caller(0)
	input, err := os.ReadFile(filepath.Join(path.Dir(filename), fmt.Sprintf("raw/%s.raw", name)))
	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

// ReplaceContent will replace blueprint placeholders with command name
func (c CommandServiceImpl) ReplaceContent(blueprint, name string) string {
	return strings.ReplaceAll(blueprint, "@@TMP@@", name)
}

// Create a new file
func (c CommandServiceImpl) Create(content, dest, fileName string) {
	path := fmt.Sprintf("%s/%s.go", helpers.GetDynamicPath(dest), strings.ToLower(fileName))
	if err := os.WriteFile(path, []byte(content), 0755); err != nil {
		log.Fatal(err)
	}
}
