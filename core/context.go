package core

import "github.com/thegogod/docs.md/core/collections"

type Context struct {
	Path string                 `json:"path"`
	Args collections.Dictionary `json:"args"`
}
