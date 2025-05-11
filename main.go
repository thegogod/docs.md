package main

import (
	"context"
	"fmt"
	"os"

	"github.com/thegogod/docs.md/commands"
	v1 "github.com/thegogod/docs.md/engines/v1"
)

func main() {
	engine := v1.New()
	context := context.WithValue(context.Background(), "engine", engine)

	if err := commands.Root.Run(context, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
