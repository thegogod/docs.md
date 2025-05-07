package core

import (
	"github.com/thegogod/docs.md/core/manifest"
	"github.com/thegogod/docs.md/markdown"
)

// determines how a project is
// constructed from a manifest
type Engine interface {
	GetPlugins() []*Plugin
	GetPlugin(name string) (*Plugin, bool)
	Render(node markdown.Node, manifest manifest.Manifest) error
}
