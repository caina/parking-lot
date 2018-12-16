package app

import "github.com/parking-lot/app/commands/help"

type Command struct {
	Action interface{}
}

func GetCommands() []Command {
	commands := make([]Command, 3)

	helpCmd := help.Instantiate()
	commands = append(commands, Command{&helpCmd})

	return commands
}
