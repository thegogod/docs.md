package build

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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
		cwd, _ := os.Getwd()
		fullpath := filepath.Join(cwd, path)
		engine := ctx.Value("engine").(core.Engine)
		manifest, err := manifest.Load(fullpath)

		if err != nil {
			return err
		}

		fmt.Println(manifest.String())
		node, err := markdown.Read(path, manifest.Build.SrcDir)

		if err != nil {
			return err
		}

		if node == nil {
			return nil
		}

		fmt.Println(node.String())

		if err = engine.Render(node, manifest); err != nil {
			return err
		}

		return nil
	},
}
