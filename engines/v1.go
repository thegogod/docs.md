package engines

import (
	"fmt"
	"html/template"
	"slices"

	"github.com/thegogod/docs.md/core"
	"github.com/thegogod/docs.md/core/manifest"
	"github.com/thegogod/docs.md/markdown"
	"github.com/thegogod/docs.md/plugins/std"
)

type V1Engine struct {
	plugins  []*core.Plugin
	template *template.Template
}

func V1(plugins ...*core.Plugin) *V1Engine {
	exists := slices.ContainsFunc(plugins, func(p *core.Plugin) bool {
		return p.Name == "std"
	})

	if !exists {
		plugins = append(plugins, std.Plugin())
	}

	for _, plugin := range plugins {
		plugin.Extend(markdown.Parser)
	}

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
	plugin.Extend(markdown.Parser)
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
