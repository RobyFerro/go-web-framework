package cli

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web-framework/tool"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// CmdCreate will create a new CLI command.
type CmdCreate struct {
	register.Command
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
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "raw/command.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", tool.GetDynamicPath("app/console"), strings.ToLower(c.Args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSUCCESS: Your %s command has been created at %s", cName, cFile)
	fmt.Printf("\nDO NOT FORGET TO REGISTER IT!")
}
