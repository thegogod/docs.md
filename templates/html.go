package templates

import "html/template"

func Html() (*template.Template, error) {
	return template.New("html").ParseFiles("html.template")
}
