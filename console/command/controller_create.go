package command

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	gwf "github.com/RobyFerro/go-web-framework"
)

// ControllerCreate will create a new controller
type ControllerCreate struct {
	Signature   string
	Description string
}

// Register this command
func (c *ControllerCreate) Register() {
	c.Signature = "controller:create <name>"
	c.Description = "Create new controller"
}

// Run this command
func (c *ControllerCreate) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	var _, filename, _, _ = runtime.Caller(0)

	cName := strings.Title(strings.ToLower(args))
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "../../raw/controller.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", gwf.GetDynamicPath("app/http/controller"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		gwf.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %sController has been created at %s", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
