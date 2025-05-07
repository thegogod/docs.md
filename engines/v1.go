package engines

import (
	"fmt"
	"html/template"
	"slices"

	"github.com/thegogod/docs.md/core"
	"github.com/thegogod/docs.md/core/manifest"
	"github.com/thegogod/docs.md/markdown"
)

type V1Engine struct {
	plugins  []*core.Plugin
	template *template.Template
}

func V1(plugins ...*core.Plugin) *V1Engine {
	return &V1Engine{plugins, template.New("main")}
}

func (self V1Engine) GetPlugins() []*core.Plugin {
	return self.plugins
}

func (self V1Engine) GetPlugin(name string) (*core.Plugin, bool) {
	i := slices.IndexFunc(self.plugins, func(p *core.Plugin) bool {
		return p.Name == name
	})

	if i == -1 {
		return nil, false
	}

	return self.plugins[i], true
}

func (self *V1Engine) AddPlugin(plugin *core.Plugin) *V1Engine {
	self.plugins = append(self.plugins, plugin)
	return self
}

func (self V1Engine) Render(node markdown.Node, manifest manifest.Manifest) error {
	for _, plugin := range self.plugins {
		if args, exists := manifest.Plugins[plugin.Name]; exists {
			if err := plugin.Import(self.template, args); err != nil {
				return err
			}
		}
	}

	fmt.Println(node.String())
	fmt.Println(manifest.String())
	return nil
}
