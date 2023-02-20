package cli

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

// CmdCreate will create a new CLI command.
type CmdCreate struct {
	entities.Command
}

// Register this command
func (c *CmdCreate) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

// Run this command
func (c *CmdCreate) Run() {
	fmt.Println("Creating new CLI command...")
	var _, filename, _, _ = runtime.Caller(0)

	splitName := strings.Split(strings.ToLower(c.Args), "_")
	for i, name := range splitName {
		splitName[i] = cases.Title(language.Und, cases.NoLower).String(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := os.ReadFile(filepath.Join(path.Dir(filename), "raw/command.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", helpers.GetDynamicPath("app/console"), strings.ToLower(c.Args))
	if err := os.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSUCCESS: Your %s command has been created at %s", cName, cFile)
	fmt.Printf("\nDO NOT FORGET TO REGISTER IT!")
}
