package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	id     int     `json:"id"`
	title  string  `json:"title"`
	artist string  `json:"artist"`
	price  float64 `json:"price"`
}

var db *sql.DB
var albums []Album

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/recordings")

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	results, errQuery := db.Query("select * from album")

	if errQuery != nil {
		log.Fatal(errQuery)
	}

	for results.Next() {
		var album Album

		results.Scan(&album.id, &album.title, &album.artist, &album.price)

		albums = append(albums, album)
	}

	fmt.Println(albums)

	results.Close()
}
