package db

import (
	"dota-api/internal"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type DB struct {
	*sqlx.DB
}

func NewDB(logger *internal.Logger) *DB {
	db, err := sqlx.Connect("mysql", os.Getenv("DATABASE_SOURCE"))
	if err != nil {
		logger.ErrorLog.Print("Can't connect to Database")
		panic(err)
	}
	return &DB{db}
}

func (d *DB) GetAllHeroes() {

}
