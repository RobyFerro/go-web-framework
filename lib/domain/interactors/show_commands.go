package interactors

import (
	"os"
	"reflect"

	"github.com/RobyFerro/go-web-framework/lib/domain/registers"
	"github.com/olekukonko/tablewriter"
)

// ShowCommands prints all available commands
type ShowCommands struct {
	Commands registers.CommandRegister
}

// Call executes show commands interactor
func (c ShowCommands) Call() {
	var data [][]string

	for _, c := range c.Commands {
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
