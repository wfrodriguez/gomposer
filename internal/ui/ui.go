package ui

import (
	"bytes"
	"text/template"
)

func Render(tpl string, values map[string]any) ([]byte, error) {
	t, err := template.New("Tags").Parse(tpl)
	if err != nil {
		return []byte{}, err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, values); err != nil {
		return []byte{}, err
	}

	return b.Bytes(), nil
}
