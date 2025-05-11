package v1

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/thegogod/docs.md/core"
	"github.com/thegogod/docs.md/core/manifest"
	"github.com/thegogod/docs.md/engines/v1/templates"
	"github.com/thegogod/docs.md/markdown"
	"github.com/thegogod/docs.md/plugins/std"
)

type Engine struct {
	plugins []*core.Plugin
}

func New(plugins ...*core.Plugin) *Engine {
	exists := slices.ContainsFunc(plugins, func(p *core.Plugin) bool {
		return p.Name == "std"
	})

	if !exists {
		plugins = append(plugins, std.Plugin())
	}

	for _, plugin := range plugins {
		plugin.Extend(markdown.Parser)
	}

	return &Engine{plugins}
}

func (self Engine) GetPlugins() []*core.Plugin {
	return self.plugins
}

func (self Engine) GetPlugin(name string) (*core.Plugin, bool) {
	i := slices.IndexFunc(self.plugins, func(p *core.Plugin) bool {
		return p.Name == name
	})

	if i == -1 {
		return nil, false
	}

	return self.plugins[i], true
}

func (self *Engine) AddPlugin(plugin *core.Plugin) *Engine {
	self.plugins = append(self.plugins, plugin)
	plugin.Extend(markdown.Parser)
	return self
}

func (self *Engine) Parse(manifest manifest.Manifest) (*template.Template, error) {
	main := template.New("main")
	fsys := os.DirFS(manifest.Build.SrcDir)
	err := fs.WalkDir(fsys, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		npath := strings.ToLower(filepath.Ext(path))

		if npath == ".md" || npath == ".html" {
			b, err := fs.ReadFile(fsys, path)

			if err != nil {
				return err
			}

			name := strings.TrimSuffix(path, filepath.Ext(path))
			_, err = main.New(name).Parse(string(b))

			if err != nil {
				return err
			}
		}

		return nil
	})

	return main, err
}

func (self Engine) Render(manifest manifest.Manifest) error {
	main, err := self.Parse(manifest)

	if err != nil {
		return err
	}

	for _, plugin := range self.plugins {
		if args, exists := manifest.Plugins[plugin.Name]; exists {
			if err := plugin.Import(main, args); err != nil {
				return err
			}
		}
	}

	var md bytes.Buffer
	err = main.ExecuteTemplate(&md, "index", manifest)

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

	content := template.Must(templates.Html.Clone())
	template.Must(content.New("index").Parse(html.String()))
	html.Reset()

	if err := content.Execute(&html, manifest); err != nil {
		return err
	}

	return os.WriteFile(
		outpath,
		html.Bytes(),
		0644,
	)
}
