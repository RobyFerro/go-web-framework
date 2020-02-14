package command

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type ModelCreate struct {
	Signature   string
	Description string
}

func (c *ModelCreate) Register() {
	c.Signature = "model:create <name>"
	c.Description = "Create new database model"
}

func (c *ModelCreate) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	var _, filename, _, _ = runtime.Caller(0)

	cName := strings.Title(strings.ToLower(args))
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "../../raw/model.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", gwf.GetDynamicPath("database/model"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		gwf.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your model %s has been created at %s", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
