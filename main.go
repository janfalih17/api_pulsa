package main

import (
	"github.com/janfalih17/api_pulsa/migrations"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	migrations.Migrate()
}
