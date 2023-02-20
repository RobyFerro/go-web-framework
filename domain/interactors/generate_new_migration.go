package interactors

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RobyFerro/go-web-framework/helpers"
)

// GenerateNewMigration generates new migrations
type GenerateNewMigration struct {
	Name string
}

// Call executes generate new migration interactor
func (c GenerateNewMigration) Call() {
	date := time.Now().Unix()
	path := helpers.GetDynamicPath("database/migration")

	filenameUp := fmt.Sprintf("%s/%d_%s.up.sql", path, date, c.Name)
	filenameDown := fmt.Sprintf("%s/%d_%s.down.sql", path, date, c.Name)

	if err := os.WriteFile(filenameUp, []byte("/* MIGRATION UP */"), 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filenameDown, []byte("/* MIGRATION DOWN */"), 0755); err != nil {
		log.Fatal(err)
	}
}
