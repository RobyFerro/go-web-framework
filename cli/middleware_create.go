package cli

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/helper"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// MiddlewareCreate will create a new http middleware
type MiddlewareCreate struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *MiddlewareCreate) Register() {
	c.Signature = "middleware:create <name>" // Change command signature
	c.Description = "Create new middleware"  // Change command description
}

// Run this command
func (c *MiddlewareCreate) Run() {
	fmt.Println("Creating new middleware...")
	var _, filename, _, _ = runtime.Caller(0)

	splitName := strings.Split(strings.ToLower(c.Args), "_")
	for i, name := range splitName {
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "raw/middleware.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", helper.GetDynamicPath("app/http/middleware"), strings.ToLower(c.Args))

	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		helper.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %s middleware has been created at %s\n", cName, cFile)
}
