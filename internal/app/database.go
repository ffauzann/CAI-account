package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	SQL SQLCred
}

func (d *Database) prepare() (err error) {
	err = d.SQL.connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	if d.SQL.AutoMigrate {
		if err = d.SQL.migrate(); err != nil {
			log.Fatal(err)
			return
		}
	}

	return
}
