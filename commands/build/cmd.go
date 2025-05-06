package build

import (
	"context"
	"fmt"

	"github.com/thegogod/docs.md/core/manifest"
	"github.com/thegogod/docs.md/markdown"
	"github.com/urfave/cli/v3"
)

var Cmd = &cli.Command{
	Name:              "build",
	Aliases:           []string{"b"},
	Description:       "build a project",
	ReadArgsFromStdin: true,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		path := cmd.Args().First()
		manifest, err := manifest.Load(path)

		if err != nil {
			return err
		}

		fmt.Println(manifest.String())
		node, err := markdown.Read(path)

		if err != nil {
			return err
		}

		if node == nil {
			return nil
		}

		fmt.Println(node.String())
		return nil
	},
}
