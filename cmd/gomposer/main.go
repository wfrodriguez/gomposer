package main

import (
	"fmt"

	"github.com/integrii/flaggy"
	"github.com/wfrodriguez/gomposer/cfg"
)

var createCommand *flaggy.Subcommand

var projectDir string = ""

var Version = "development"

func flagSettings() {
	createCommand = flaggy.NewSubcommand("new")
	createCommand.Description = "Crea un nuevo proyecto."
	createCommand.AddPositionalValue(&projectDir, "projectDir", 1, true, "Nombre del proyecto.")
	flaggy.AttachSubcommand(createCommand, 1)

	flaggy.SetVersion(Version)
}

func main() {
	fmt.Println(cfg.Logo)
	flaggy.Parse()
}
