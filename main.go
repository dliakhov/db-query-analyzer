package main

import (
	"fmt"
	"github.com/dliakhov/db-query-analyzer/cmd"
	"github.com/pkg/errors"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	cli := cmd.NewCLI()
	//cli.Version = fmt.Sprintf("%s (Commit: %s)", version, commit)

	if err := cli.Execute(); err != nil {
		return errors.Wrap(err, "error initializing the command")
	}
	return nil
}
