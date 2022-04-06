package db

import (
	"dota-api/internal"
	"dota-api/internal/models"
	"fmt"
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

func (d *DB) GetAllHeroes() *[]models.Hero {
	var hero []models.Hero
	err := d.DB.Select(&hero, "SELECT \n\th.name as hero_name,\n    h.localized_name as localized_name,\n    c.name as class_name,\n    h.image_url\nFROM\n\thero h\nINNER JOIN\n\tclass c ON c.id = h.class_id")
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return &hero
}
