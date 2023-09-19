package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/integrii/flaggy"
	"github.com/wfrodriguez/gomposer/cfg"
	"github.com/wfrodriguez/gomposer/internal/command"
)

var createCommand *flaggy.Subcommand
var indexCommand *flaggy.Subcommand

var projectDir string = ""

var Version = "development"

func flagSettings() {
	createCommand = flaggy.NewSubcommand("new")
	createCommand.Description = "Crea un nuevo proyecto."
	createCommand.AddPositionalValue(&projectDir, "projectDir", 1, true, "Nombre del proyecto.")
	flaggy.AttachSubcommand(createCommand, 1)

	indexCommand = flaggy.NewSubcommand("index")
	indexCommand.Description = "Genera los archivos tags e índice principal del proyecto"
	flaggy.AttachSubcommand(indexCommand, 1)

	flaggy.SetVersion(Version)
	flaggy.SetName("Gomposer")
}

func main() {
	fmt.Println(cfg.Logo)
	var err error

	projectDir, err = os.Getwd()
	if err != nil {
		panic(err) // TODO: handle error
	}
	// FIXME: Para pruebas de desarrollo
	projectDir = filepath.Join(projectDir, "__test")

	flagSettings()

	flaggy.Parse()

	if createCommand.Used {
	} else if indexCommand.Used {
		command.IndexProject(projectDir)
	}
}
