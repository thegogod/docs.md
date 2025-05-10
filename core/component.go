package core

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"maps"

	"github.com/yuin/goldmark"
)

type Component struct {
	Name      string                                    `json:"name"`
	Template  string                                    `json:"template,omitempty"`
	Assets    []string                                  `json:"assets,omitempty"`
	Functions map[string]func(args ...any) (any, error) `json:"-"`

	template *template.Template
}

func (self Component) Select(tag string) bool {
	return self.Name == tag
}

func (self *Component) Extend(markdown goldmark.Markdown) {
}

func (self *Component) Import(plugin string, parent *template.Template) (*template.Template, error) {
	if self.template != nil {
		return self.template, nil
	}

	funcs := template.FuncMap{}

	for name, fn := range self.Functions {
		funcs[name] = fn
	}

	template, err := parent.New(fmt.Sprintf("%s->%s", plugin, self.Name)).
		Funcs(funcs).
		Parse(self.Template)

	self.template = template
	return template, err
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
