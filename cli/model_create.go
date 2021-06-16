package cli

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/tool"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// ModelCreate will create a new Gorm model
type ModelCreate struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *ModelCreate) Register() {
	c.Signature = "model:create <name>"
	c.Description = "Create new database model"
}

// Run this command
func (c *ModelCreate) Run() {
	fmt.Println("Creating new model...")
	var _, filename, _, _ = runtime.Caller(0)

	cName := strings.Title(strings.ToLower(c.Args))
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "raw/model.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", tool.GetDynamicPath("database/model"), strings.ToLower(c.Args))

	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("(model\\.[A-Za-z]\\w+( *){},(\\n*)(\\t*| *))")
	newModel := fmt.Sprintf("model.%s{},\n\t\t", cName)
	autoRegister(re, newModel)

	fmt.Printf("\nSUCCESS: Your model %s has been created at %s", cName, cFile)
}
