package gwf

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// CmdCreate will create a new CLI command.
type CmdCreate struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *CmdCreate) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

// Run this command
func (c *CmdCreate) Run() {
	var _, filename, _, _ = runtime.Caller(0)

	splitName := strings.Split(strings.ToLower(c.Args), "_")
	for i, name := range splitName {
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "raw/command.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", GetDynamicPath("app/console/command"), strings.ToLower(c.Args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %s command has been created at %s\n", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
