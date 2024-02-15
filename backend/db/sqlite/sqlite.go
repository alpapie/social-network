package sqlite

import (
	"fmt"
	"log"
	controller "social_network/controller"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func ApplyMigrations() error {
	driver, err := sqlite.WithInstance(controller.DB, &sqlite.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/sqlite",
		"sqlite3", driver)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return err
	}

	return nil
}

func RollbackMigrations() error {
	driver, err := sqlite.WithInstance(controller.DB, &sqlite.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://..db/migrations/sqlite",
		"sqlite3", driver)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
		return err
	}

	return nil
}
