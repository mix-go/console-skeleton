package commands

import (
    "github.com/mix-go/console-skeleton/commands"
    "github.com/mix-go/console"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "db",
            Usage: "\tDatabase query demo",
            Command: &commands.DBCommand{},
        },
    )
}
