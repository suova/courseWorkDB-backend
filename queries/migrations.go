package queries

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func doMigrate() {

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}