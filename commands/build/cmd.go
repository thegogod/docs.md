package build

import (
	"context"

	"github.com/thegogod/docs.md/core"
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
		engine := ctx.Value("engine").(core.Engine)
		manifest, err := manifest.Load(path)

		if err != nil {
			return err
		}

		node, err := markdown.Read(path)

		if err != nil {
			return err
		}

		if node == nil {
			return nil
		}

		if err = engine.Render(node, manifest); err != nil {
			return err
		}

		return nil
	},
}
