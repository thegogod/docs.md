package core

import (
	"bytes"
	"errors"
	"html/template"
	"maps"

	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
)

var KindComponent = gast.NewNodeKind("Component")

type Component struct {
	gast.BaseInline

	Selector  string                                    `json:"selector"`
	Template  string                                    `json:"template,omitempty"`
	Assets    []string                                  `json:"assets,omitempty"`
	Functions map[string]func(args ...any) (any, error) `json:"-"`

	template *template.Template
}

func (self Component) Select(tag string) bool {
	return self.Selector == tag
}

func (self *Component) Extend(markdown goldmark.Markdown) {
}

func (self *Component) Import(parent *template.Template) error {
	funcs := template.FuncMap{}

	for name, fn := range self.Functions {
		funcs[name] = fn
	}

	template, err := parent.New(self.Selector).
		Funcs(funcs).
		Parse(self.Template)

	self.template = template
	return err
}

func (self Component) Render(context Context) ([]byte, error) {
	if self.template == nil {
		return nil, errors.New("component must be imported before use")
	}

	var buffer bytes.Buffer
	data := maps.Clone(context.Args).Set("assets", self.Assets)

	if err := self.template.Execute(&buffer, data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// func (self *Component) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {

// }

// func (self *Component) Kind() gast.NodeKind {
// 	return KindComponent
// }

// func (self *Component) Dump(source []byte, level int) {
// 	data := map[string]string{
// 		"selector": self.Selector,
// 		"template": self.Template,
// 		"assets":   strings.Join(self.Assets, ", "),
// 	}

// 	gast.DumpHelper(self, source, level, data, nil)
// }
