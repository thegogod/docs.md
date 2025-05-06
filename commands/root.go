package commands

import (
	"context"
	"fmt"

	"github.com/thegogod/docs.md/commands/build"
	"github.com/urfave/cli/v3"
)

var Root = &cli.Command{
	Name:        "docs.md",
	Description: "generate a documentation website from markdown",
	Commands:    []*cli.Command{build.Cmd},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		fmt.Println("root", cmd.Args())
		return nil
	},
}
