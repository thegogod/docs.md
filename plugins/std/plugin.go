package std

import (
	"fmt"

	"github.com/thegogod/docs.md/core"
	"github.com/thegogod/docs.md/core/collections"
)

func Plugin() *core.Plugin {
	return &core.Plugin{
		Name:       "std",
		Version:    "0.0.0",
		Components: []*core.Component{Button},
		OnInit: func(args collections.Dictionary) error {
			fmt.Println("initialized")
			return nil
		},
	}
}
