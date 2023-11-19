package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConectarBancoDeDados(){
	conn,err := sql.Open("mysql",os.Getenv("CONNECTION"))

	if err != nil {
		log.Fatal(err.Error())
	}
	
	log.Println("Banco de dados incializado!")

	DB = conn
}