package markdown

import "fmt"

type Node interface {
	fmt.Stringer

	GetName() string
	GetNodes() []Node
}
