package main

import (
	"context"
	"fmt"
	"os"

	"github.com/thegogod/docs.md/commands"
	"github.com/thegogod/docs.md/engines"
)

func main() {
	engine := engines.V1()
	context := context.WithValue(context.Background(), "engine", engine)

	if err := commands.Root.Run(context, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
