package files

import (
	"text/template"
	"bytes"
)


type Template struct {
	tpl *template.Template
}


func ParseTemplate(base string) (*Template, error) {
	tpl, err := template.New("base").Parse(base)
	if err != nil {
		return nil, err
	}

	return &Template {
		tpl: tpl,
	}, nil
}


func (template *Template) Template(obj interface{}) (string, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := template.tpl.Execute(buffer, obj)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
