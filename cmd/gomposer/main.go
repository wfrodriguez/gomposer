package main

import (
	"fmt"
	"os"

	"github.com/integrii/flaggy"
	"github.com/wfrodriguez/console"
	"github.com/wfrodriguez/gomposer/cfg"
	"github.com/wfrodriguez/gomposer/internal/command"
)

var createCommand *flaggy.Subcommand
var indexCommand *flaggy.Subcommand

var projectDir string = ""
var projectName string = ""

var Version = "development"

func flagSettings() {
	createCommand = flaggy.NewSubcommand("new")
	createCommand.Description = "Crea un nuevo proyecto."
	createCommand.AddPositionalValue(&projectName, "projectName", 1, true, "Nombre del proyecto.")
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
	// projectDir = filepath.Join(projectDir, "__test")

	flagSettings()

	flaggy.Parse()

	if createCommand.Used {
		command.NewProject(projectDir, projectName)
	} else if indexCommand.Used {
		command.IndexProject(projectDir)
	} else {
		console.Error("No se especificó ninguna opción")
		os.Exit(10)
	}
}
