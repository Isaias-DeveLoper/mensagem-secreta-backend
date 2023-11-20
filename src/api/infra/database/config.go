package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConectarBancoDeDados() {

	var drive string

	if os.Getenv("ENV") == "development" {
		drive = "mysql"
	} else {
		drive = "postgres"
	}

	conn, err := sql.Open(drive, os.Getenv("CONNECTION"))

	if err != nil {
		log.Fatal(err.Error())
	}

	if os.Getenv("ENV") != "development" {
		
		driver, err := postgres.WithInstance(conn, &postgres.Config{})

		if os.Getenv("ENV") == "development" {
			driver,err = mysql.WithInstance(conn,&mysql.Config{})
		} 

		if err != nil {
			panic(err.Error())
		}
		m, e:= migrate.NewWithDatabaseInstance(
			"file://src/api/infra/database/migrations",
			drive, driver)
		
		if e != nil {
			panic(e.Error())
		}
		
		m.Up()

	}

	log.Println("Banco de dados incializado!")

	DB = conn
}
