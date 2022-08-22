package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web-framework/tool"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ControllerCreate will create a new controller
type ControllerCreate struct {
	register.Command
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

	cName := cases.Title(language.Und, cases.NoLower).String(c.Args)
	input, _ := os.ReadFile(filepath.Join(path.Dir(filename), "raw/controller.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", tool.GetDynamicPath("app/http/controller"), strings.ToLower(c.Args))
	if err := os.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(&controller\\.[A-Za-z]\\w+( *){},(\\n*)(\\t*| *))`)
	newController := fmt.Sprintf("\t&controller.%sController{},\n\t\t", cName)
	autoRegister(re, newController)

	fmt.Printf("\nSUCCESS: Your %sController has been created at %s", cName, cFile)
}
