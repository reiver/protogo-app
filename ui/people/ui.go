package peopleui

import (
	_ "embed"
	htmltemplate "html/template"
	"strings"

	"github.com/reiver/go-erorr"
)

const (
	errNilTemplate = erorr.Error("nil template")
)

//go:embed template.html
var tmpl string

var template *htmltemplate.Template = htmltemplate.Must(htmltemplate.New("peopleui").Parse(tmpl))

type model struct {
	People []Person
}

type Person struct {
	Address string
	AvatarURI htmltemplate.HTML
	Name string
}

func Render(people []Person) (string, error) {
	if nil == template {
		var nada string
		return nada, errNilTemplate
	}

	var builder strings.Builder

	data := model{
		People: people,
	}

	err := template.Execute(&builder, data)
	if nil != err {
		var nada string
		return nada, err
	}

	return builder.String(), nil
}
