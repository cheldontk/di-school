package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/cheldontk/di-school/di-go/application/config"
)

var DB *sql.DB

func Connect() error {
	var err error
	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parse str to int")
	}

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	CreateProductTable()
	fmt.Println("Connection Opened to Database")

	return nil
}
