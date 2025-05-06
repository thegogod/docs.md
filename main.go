package main

import (
	"context"
	"fmt"
	"os"

	"github.com/thegogod/docs.md/commands"
)

func main() {
	if err := commands.Root.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
