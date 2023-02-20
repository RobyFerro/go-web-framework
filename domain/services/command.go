package services

// CommandService defines methods used to generate new CLI commands
type CommandService interface {
	SnakeToCamelCase(raw string) string

	ReadBlueprint(name string) string

	ReplaceContent(blueprint, name string) string

	Create(content, dest, fileName string)
}
