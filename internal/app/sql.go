package app

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

type SQLCred struct {
	Driver      string
	Host        string
	Port        uint32
	Schema      string
	Username    string
	Password    string
	AutoMigrate bool
	DB          *sqlx.DB
}

func (d *SQLCred) connect() (err error) {
	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		d.Driver,
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Schema,
	)

	d.DB, err = sqlx.Connect(d.Driver, dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func (d *SQLCred) migrate() (err error) {
	driver, err := postgres.WithInstance(d.DB.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/migration/", d.Driver, driver)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = m.Up(); err != nil {
		// Bypass err no change or no changes
		if err == migrate.ErrNoChange {
			return nil
		}

		// Bypass err no change or no migrations
		if e, ok := err.(*fs.PathError); ok {
			if e.Err == fs.ErrNotExist {
				return nil
			}
		}

		log.Fatal(err)
		return
	}

	return
}
