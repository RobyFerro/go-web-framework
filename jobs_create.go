package gwf

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// JobCreate will create a new async job
type JobCreate struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *JobCreate) Register() {
	c.Signature = "job:create <name>"      // Change command signature
	c.Description = "Create new async job" // Change command description
}

// Run this command
func (c *JobCreate) Run() {
	fmt.Println("Creating new async jobs...")
	var _, filename, _, _ = runtime.Caller(0)

	splitName := strings.Split(strings.ToLower(c.Args), "_")
	for i, name := range splitName {
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(filepath.Join(path.Dir(filename), "raw/job.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", GetDynamicPath("job"), strings.ToLower(c.Args))

	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %s job has been created at %s\n", cName, cFile)
}
