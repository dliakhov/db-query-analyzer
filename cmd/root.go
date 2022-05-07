package cmd

import (
	"github.com/dliakhov/db-query-analyzer/config"
	_ "github.com/jackc/pgx/v4"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/cobra"
	"log"
)

var conf *config.Config

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use: "db-query-analyzer",
	}

	cli.AddCommand(newHttpService())

	conf = &config.Config{}
	_, err := flags.Parse(conf)
	if err != nil {
		log.Fatalf("Error during parsing configuration: %v", err)
	}

	return cli
}
