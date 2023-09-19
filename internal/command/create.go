package command

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
//
// 	"github.com/tidwall/sjson"
// 	"github.com/wfrodriguez/gomposer/internal"
// )
//
// func CreateCommand() {
// 	validate := func(input string) error {
// 		if input == "" {
// 			return fmt.Errorf("el campo es obligatorio")
// 		}
// 		return nil
// 	}
//
// 	value, err := internal.Prompt("Nombre del proyecto", validate)
// 	internal.handleError(err, "Error al obtener el nombre del proyecto")
// 	slug := internal.Slugify(value)
// 	wd := internal.WorkingDir()
// 	internal.SimpleLogger().Info("Creando proyecto `%s`...", slug)
// 	os.MkdirAll(filepath.Join(wd, slug), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `post`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "post"), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `template`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "template"), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `static`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "static"), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `static/img`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "static", "img"), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `static/js`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "static", "js"), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `static/css`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "static", "css"), 0755)
// 	internal.SimpleLogger().Info("Creando carpeta `static/font`...")
// 	os.MkdirAll(filepath.Join(wd, slug, "static", "font"), 0755)
// 	err = generateConfig(value)
// 	internal.HandleError(err, "Error al generar el archivo de configuración")
// }
//
// func generateConfig(name string) error {
// 	data := `{}`
// 	data = sjson.Set(data, "config.title", name)
// 	data = sjson.Set(data, "vars.name", "value")
//
// 	return os.WriteFile("variables.json", []byte(data), 0644)
// }
