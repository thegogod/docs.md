package core

import "github.com/thegogod/docs.md/core/collections"

// an integration that can augment
// how the engine functions
type Plugin interface {
	GetName() string
	GetVersion() string

	Configure(config collections.Dictionary) error
	Render(context Context) ([]byte, error)
}
