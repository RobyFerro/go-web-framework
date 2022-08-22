package cli

import (
	"log"
	"os"
	"regexp"
	"strings"
)

// Handles Go-Web component auto Register
func autoRegister(re *regexp.Regexp, newElement string) {
	var elements []string
	content, err := os.ReadFile("register/register.go")
	if err != nil {
		log.Fatal(err)
	}

	res := re.FindAll(content, -1)

	for _, cN := range res {
		elements = append(elements, string(cN))
	}

	oldString := strings.Join(elements, "")
	elements = append(elements, newElement)
	newString := strings.Join(elements, "")

	newContent := strings.Replace(string(content), oldString, newString, 1)
	if err := os.WriteFile("register/register.go", []byte(newContent), 0644); err != nil {
		log.Fatal(err)
	}
}
