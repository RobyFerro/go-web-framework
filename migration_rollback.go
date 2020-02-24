package gwf

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// MigrateRollback will rollback some migration in your database
type MigrateRollback struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *MigrateRollback) Register() {
	c.Signature = "migration:rollback <steps>"
	c.Description = "Rollback migrations"
}

// Run this command
func (c *MigrateRollback) Run(db *gorm.DB) {
	step, _ := strconv.Atoi(c.Args)
	batch := getLastBatch(db)

	for i := 0; i < step; i++ {
		var migrations []migration
		if err := db.Order("created_at", true).Where("batch LIKE ?", batch).Find(&migrations).Error; err != nil {
			ProcessError(err)
		}

		// Execute given rollback
		rollbackMigrations(migrations, db)
		batch--
	}
}

// Core of rollback method.
// This method will parse a given set of migration and run the relative rollback
func rollbackMigrations(migrations []migration, db *gorm.DB) {
	for _, m := range migrations {
		rollbackFile := strings.ReplaceAll(m.Name, ".up.sql", ".down.sql")
		fmt.Printf("\nRolling back '%s' migration...\n", rollbackFile)

		if payload, err := ioutil.ReadFile(rollbackFile); err != nil {
			ProcessError(err)
		} else {
			db.Exec(string(payload)).Row()
		}

		if err := db.Unscoped().Delete(&m).Error; err != nil {
			ProcessError(err)
		}

		fmt.Printf("Success! %s has been rolled back!", rollbackFile)
	}
}
