package gwf

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// ControllerCreate will create a new controller
type ControllerCreate struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *ControllerCreate) Register() {
	c.Signature = "controller:create <name>"
	c.Description = "Create new controller"
}

// Run this command
func (c *ControllerCreate) Run() {
	fmt.Println("Creating new controller...")
	var _, filename, _, _ = runtime.Caller(0)

	cName := strings.Title(strings.ToLower(c.Args))
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "raw/controller.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", GetDynamicPath("app/http/controller"), strings.ToLower(c.Args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		ProcessError(err)
	}

	re := regexp.MustCompile("(&controller\\.[A-Za-z]\\w+( *){},(\\n*)(\\t*| *))")
	newController := fmt.Sprintf("\t&controller.%sController{},\n\t\t", cName)
	autoRegister(re, newController)

	fmt.Printf("\nSUCCESS: Your %sController has been created at %s", cName, cFile)
}
