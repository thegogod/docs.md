package markdown

import (
	"fmt"
	"html/template"
)

type Node interface {
	fmt.Stringer

	GetPath() string
	GetRelPath() string
	GetName() string
	GetNodes() []Node

	Parse(template *template.Template) (*template.Template, error)
}
