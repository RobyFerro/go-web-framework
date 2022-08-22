package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web-framework/tool"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// MiddlewareCreate will create a new http middleware
type MiddlewareCreate struct {
	register.Command
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
		splitName[i] = cases.Title(language.Und, cases.NoLower).String(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := os.ReadFile(filepath.Join(path.Dir(filename), "raw/middleware.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", tool.GetDynamicPath("app/http/middleware"), strings.ToLower(c.Args))

	if err := os.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSUCCESS: Your %s middleware has been created at %s\n", cName, cFile)
	fmt.Println("Do not forget to register your new middleware!")
}
