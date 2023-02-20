package cli

import (
	"os"
	"reflect"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/registers"
	"github.com/olekukonko/tablewriter"
)

// ShowCommands will show all registered commands
type ShowCommands struct {
	entities.Command
}

// Register this command
func (c *ShowCommands) Register() {
	c.Signature = "show:commands"
	c.Description = "Show Go-Web commands list"
}

// Run this command
func (c *ShowCommands) Run(commands registers.CommandRegister) {

	var data [][]string

	for _, c := range commands {
		m := reflect.ValueOf(c).MethodByName("Register")
		m.Call([]reflect.Value{})

		cmd := reflect.ValueOf(c).Elem()

		signature := cmd.FieldByName("Signature").String()
		description := cmd.FieldByName("Description").String()

		data = append(data, []string{signature, description})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"COMMAND", "DESCRIPTION"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
