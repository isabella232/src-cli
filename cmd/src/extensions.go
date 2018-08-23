package main

import (
	"flag"
	"fmt"
)

var extensionsCommands commander

func init() {
	usage := `'src extensions' is a tool that manages extensions in the extension registry on a Sourcegraph instance.

Usage:

	src extensions command [command options]

The commands are:

	list       lists extensions
	get        gets an extension

Use "src extensions [command] -h" for more information about a command.

Alias: "src ext"
`

	flagSet := flag.NewFlagSet("extensions", flag.ExitOnError)
	handler := func(args []string) error {
		extensionsCommands.run(flagSet, "src extensions", usage, args)
		return nil
	}

	// Register the command.
	commands = append(commands, &command{
		flagSet: flagSet,
		aliases: []string{"ext", "extension"},
		handler: handler,
		usageFunc: func() {
			fmt.Println(usage)
		},
	})
}

const registryExtensionFragment = `
fragment RegistryExtensionFields on RegistryExtension {
    id
    uuid
    extensionID
    name
    createdAt
    updatedAt
    url
    remoteURL
    registryName
    isLocal
    manifest {
        raw
        title
        description
    }
}
`

type Extension struct {
	ID           string
	UUID         string
	ExtensionID  string
	Name         string
	CreatedAt    string
	UpdatedAt    string
	URL          string
	RemoteURL    string
	RegistryName string
	IsLocal      bool
	Manifest     struct {
		Raw         string
		Title       string
		Description string
	}
}
