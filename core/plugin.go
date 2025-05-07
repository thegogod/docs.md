package core

import (
	"html/template"

	"github.com/thegogod/docs.md/core/collections"
)

// an integration that can augment
// how the engine functions
type Plugin struct {
	Name       string                                  `json:"name"`
	Version    string                                  `json:"string,omitempty"`
	Components []Component                             `json:"components,omitempty"`
	OnInit     func(args collections.Dictionary) error `json:"-"`

	template *template.Template
}

func (self *Plugin) Import(template *template.Template, args collections.Dictionary) error {
	if self.template != nil {
		return nil
	}

	template = template.New(self.Name)

	for _, component := range self.Components {
		if err := component.Import(template); err != nil {
			return err
		}
	}

	if self.OnInit != nil {
		if err := self.OnInit(args); err != nil {
			return err
		}
	}

	self.template = template
	return nil
}

func (self Plugin) Select(tag string) (Component, bool) {
	for _, component := range self.Components {
		if match := component.Select(tag); match {
			return component, true
		}
	}

	return Component{}, false
}
