package cmd

import (
	"fmt"
	"github.com/dliakhov/db-query-analyzer/internal/httpservice"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newHttpService() *cobra.Command {
	httpServer := &cobra.Command{
		Use:   "httpservice",
		Short: "Starts HTTP server",
		RunE:  runHttpService,
	}
	return httpServer
}

func runHttpService(cmd *cobra.Command, args []string) error {
	log, _ := zap.NewDevelopment()

	dbConf := conf.Database
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbConf.Username, dbConf.Password, dbConf.Hostname, dbConf.Port, dbConf.Name))
	if err != nil {
		log.Error(fmt.Sprintf("cannot connect to databse: %v", err))
		return err
	}
	defer db.Close()

	err = httpservice.Run(conf, db)
	if err != nil {
		log.Error(fmt.Sprintf("error during running server: %v", err))
		return err
	}
	return nil
}
