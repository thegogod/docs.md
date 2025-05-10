package engines

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"slices"

	"github.com/thegogod/docs.md/core"
	"github.com/thegogod/docs.md/core/manifest"
	"github.com/thegogod/docs.md/markdown"
	"github.com/thegogod/docs.md/plugins/std"
)

type V1Engine struct {
	plugins []*core.Plugin
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

	return &V1Engine{plugins}
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

func (self *V1Engine) Parse(manifest manifest.Manifest) (*template.Template, error) {
	return template.ParseGlob(filepath.Join(manifest.Build.SrcDir, "*.md"))
}

func (self V1Engine) Render(manifest manifest.Manifest) error {
	template, err := self.Parse(manifest)

	if err != nil {
		return err
	}

	for _, plugin := range self.plugins {
		if args, exists := manifest.Plugins[plugin.Name]; exists {
			if err := plugin.Import(template, args); err != nil {
				return err
			}
		}
	}

	var md bytes.Buffer
	err = template.ExecuteTemplate(&md, "index.md", map[string]any{})

	if err != nil {
		return err
	}

	var html bytes.Buffer

	if err := markdown.Parser.Convert(md.Bytes(), &html); err != nil {
		return err
	}

	outpath := filepath.Join(manifest.Build.OutDir, "index.html")
	fmt.Printf("%s => writing...\n", outpath)

	if _, err := os.Stat(filepath.Dir(outpath)); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}

		if err := os.MkdirAll(filepath.Dir(outpath), 0777); err != nil {
			return err
		}
	}

	return os.WriteFile(
		outpath,
		html.Bytes(),
		0644,
	)
}
