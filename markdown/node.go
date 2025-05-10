package markdown

import "fmt"

type Node interface {
	fmt.Stringer

	GetPath() string
	GetRelPath() string
	GetName() string
	GetNodes() []Node
}
