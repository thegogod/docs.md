package core

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"maps"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Component struct {
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
	fmt.Println("extend...")
	markdown.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(nil, 200),
	))
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
