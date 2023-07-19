package internal

import (
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/gosimple/slug"
	"github.com/manifoldco/promptui"
)

// Funciones de utilidad varias

// TernaryIf es una función que devuelve un valor dependiendo de una condición, similar al operador ternario `?:`
// para que funcione `T` debe ser del mismo tipo de dato
func TernaryIf[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Slugify genera una cadena slug a partir de una cadena Unicode, URL-amigable con soporte para múltiples idiomas.
func Slugify(s string) string {
	return slug.Make(s)
}

func Prompt(label string, valid func(string) error) (value string, err error) {

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green | bold }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  valid,
	}

	value, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return
}

// WorkingDir devuelve el directorio de trabajo actual
func WorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		SimpleLogger().Error("Error obteniendo ubicación del directorio de trabajo activo: %s", err.Error())
		os.Exit(3)
	}

	return dir
}

// RuneCount devuelve el total de caracteres existentes en un string a diferencia de `len()` que lee la cantidad de
// bytes de un elemento
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

func findFileByExt(path, ext string) ([]string, error) {
	var files = make([]string, 0)

	err := filepath.Walk(path, func(p string, in os.FileInfo, er error) error {
		if er != nil {
			return er
		}

		if !in.IsDir() && filepath.Ext(p) == "."+ext {
			files = append(files, p)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
