package main

import (
	"context"
	"fmt"
	"go-micro/cmd/migrations/migrations"
	"go-micro/internal/pkg/db"
	"log"
	"os"
	"strings"

	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "manage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "environment",
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name: "db",
				Subcommands: []*cli.Command{
					{
						Name:  "init",
						Usage: "create migration tables",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)
							return migrator.Init(context.Background())
						},
					},
					{
						Name:  "migrate",
						Usage: "migrate database",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)

							group, err := migrator.Migrate(context.Background())
							if err != nil {
								return err
							}

							if group.ID == 0 {
								fmt.Printf("there are no new migrations.Migrations to run\n")
								return nil
							}

							fmt.Printf("migrated to %s\n", group)
							return nil
						},
					},
					{
						Name:  "rollback",
						Usage: "rollback the last migration group",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)

							group, err := migrator.Rollback(context.Background())
							if err != nil {
								return err
							}

							if group.ID == 0 {
								fmt.Printf("there are no groups to roll back\n")
								return nil
							}

							fmt.Printf("rolled back %s\n", group)
							return nil
						},
					},
					{
						Name:  "lock",
						Usage: "lock migrations.Migrations",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)
							return migrator.Lock(context.Background())
						},
					},
					{
						Name:  "unlock",
						Usage: "unlock migrations.Migrations",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)
							return migrator.Unlock(context.Background())
						},
					},
					{
						Name:  "create_go",
						Usage: "create Go migration",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)

							name := strings.Join(c.Args().Slice(), "_")
							mf, err := migrator.CreateGoMigration(context.Background(), name)
							if err != nil {
								return err
							}
							fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)

							return nil
						},
					},
					{
						Name:  "create_sql",
						Usage: "create up and down SQL migrations.Migrations",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)

							name := strings.Join(c.Args().Slice(), "_")
							files, err := migrator.CreateSQLMigrations(context.Background(), name)
							if err != nil {
								return err
							}

							for _, mf := range files {
								fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
							}

							return nil
						},
					},
					{
						Name:  "status",
						Usage: "print migrations.Migrations status",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)

							ms, err := migrator.MigrationsWithStatus(context.Background())
							if err != nil {
								return err
							}
							fmt.Printf("migrations.Migrations: %s\n", ms)
							fmt.Printf("unapplied migrations.Migrations: %s\n", ms.Unapplied())
							fmt.Printf("last migration group: %s\n", ms.LastGroup())

							return nil
						},
					},
					{
						Name:  "mark_applied",
						Usage: "mark migrations.Migrations as applied without actually running them",
						Action: func(c *cli.Context) error {
							migrator := migrate.NewMigrator(db.DB(), migrations.Migrations)

							group, err := migrator.Migrate(context.Background(), migrate.WithNopMigration())
							if err != nil {
								return err
							}

							if group.ID == 0 {
								fmt.Printf("there are no new migrations.Migrations to mark as applied\n")
								return nil
							}

							fmt.Printf("marked as applied %s\n", group)
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
