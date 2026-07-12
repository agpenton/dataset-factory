package prompt

import (
	"bytes"
	"text/template"
)

type Renderer struct{}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(
	templatePath string,
	data TemplateData,
) (string, error) {

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer

	if err := tpl.Execute(&out, data); err != nil {
		return "", err
	}

	return out.String(), nil
}
