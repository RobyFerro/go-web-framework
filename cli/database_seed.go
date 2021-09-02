package cli

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/register"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

// Seeder will handle database seeding.
type Seeder struct {
	register.Command
}

// Register this command
func (c *Seeder) Register() {
	c.Signature = "database:seed <name>"
	c.Description = "Execute database seeder"
}

// Run this command
// Todo: Improve this method to run a single seeder
func (c *Seeder) Run(db *gorm.DB, models register.ModelRegister) {
	fmt.Println("Execute database seeding...")
	if len(c.Args) > 0 {
		extractSpecificModel(c.Args, &models.List)
	}

	seed(models.List, db)
}

// Extract the specified models from model list
func extractSpecificModel(name string, models *[]interface{}) {
	var newModels []interface{}

	for _, m := range *models {
		modelName := reflect.TypeOf(m).Name()

		if strings.EqualFold(name, modelName) {
			newModels = append(newModels, m)
			break
		}
	}

	*models = newModels
}

// Parse model register and run every seed
func seed(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		fmt.Printf("\nCreating items for model %s...\n", reflect.TypeOf(m).Name())
		v := reflect.ValueOf(m)
		method := v.MethodByName("Seed")
		method.Call([]reflect.Value{reflect.ValueOf(db)})

		fmt.Printf("Success!\n")
	}

	fmt.Println("Seeding complete!")
}
