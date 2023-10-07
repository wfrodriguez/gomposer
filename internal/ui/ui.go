package ui

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"
	"time"
)

//go:embed "all:skeleton"
var Skeleton embed.FS

var monthEsp = map[string]string{
	"01": "Enero",
	"02": "Febrero",
	"03": "Marzo",
	"04": "Abril",
	"05": "Mayo",
	"06": "Junio",
	"07": "Julio",
	"08": "Agosto",
	"09": "Septiembre",
	"10": "Octubre",
	"11": "Noviembre",
	"12": "Diciembre",
}
var funcs = template.FuncMap{
	"now": func() string {
		data := strings.Split(time.Now().Format("01:02:2006"), ":")
		data[0] = monthEsp[data[0]]
		return fmt.Sprintf("%s %s %s", data[1], data[0], data[2])
	},
}

func Render(tpl string, values map[string]any) ([]byte, error) {
	t, err := template.New("Tags").Funcs(funcs).Parse(tpl)
	if err != nil {
		return []byte{}, err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, values); err != nil {
		return []byte{}, err
	}

	return b.Bytes(), nil
}
