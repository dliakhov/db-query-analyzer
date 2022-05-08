package cmd

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const migrationsFilePath = "file://migrations"

func newMigrateUpCommand() *cobra.Command {
	dbConf := conf.Database

	return &cobra.Command{
		Use:   "migrate",
		Short: "Executes database migrations.",
		RunE: func(cmd *cobra.Command, args []string) error {
			log, _ := zap.NewDevelopment()

			db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				dbConf.Username, dbConf.Password, dbConf.Hostname, dbConf.Port, dbConf.Name))
			if err != nil {
				log.Error("cannot connect to database", zap.Error(err))
				return err
			}

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				log.Error("cannot get driver to database", zap.Error(err))
				return err
			}

			m, err := migrate.NewWithDatabaseInstance(migrationsFilePath, "postgres", driver)
			if err != nil {
				log.Error("cannot create migrate instance", zap.Error(err))
				return err
			}

			if err := m.Up(); err != nil {
				if err == migrate.ErrNoChange {
					log.Info("No changes")
					return nil
				}
				return err

			}
			return nil
		},
	}
}
