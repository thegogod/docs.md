package core

import (
	"github.com/thegogod/docs.md/core/manifest"
)

// determines how a project is
// constructed from a manifest
type Engine interface {
	GetPlugins() []*Plugin
	GetPlugin(name string) (*Plugin, bool)
	Render(manifest manifest.Manifest) error
}
