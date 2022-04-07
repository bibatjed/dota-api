package db

import (
	"dota-api/internal"
	"dota-api/internal/models"
	"dota-api/internal/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
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

func (d *DB) GetAllHeroes(pagination utils.Pagination) []models.Hero {
	var hero []models.Hero
	query := `
	SELECT 
		h.name as hero_name,
    	h.localized_name as localized_name,
    	c.name as class_name,
    	h.image_url
	FROM
		hero h
	INNER JOIN
		class c ON c.id = h.class_id`

	query += "\nLIMIT " + strconv.Itoa(pagination.Limit)
	query += " OFFSET " + strconv.Itoa(pagination.Offset)

	fmt.Println(query)

	err := d.DB.Select(&hero, query)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	if len(hero) < 1 {
		return make([]models.Hero, 0)
	}
	return hero
}

func (d *DB) GetAllHeroesCount() int {
	query := `
	SELECT 
		count(*) as hero_count
	FROM
		hero h`


	var count int

	err := d.DB.QueryRow(query).Scan(&count)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	return count
}
