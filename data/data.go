/*
Copyright © 2022 Junjie Yang junjie@jyang.dev
*/
package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDataBase() error {
	var err error
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return err
	}
	os.MkdirAll(fmt.Sprintf("%s/.config/nav-go", home), 0755)
	db, err = sql.Open("sqlite3", fmt.Sprintf("%s/.config/nav-go/nav-database.db", home))

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS navgo (
    "name" TEXT NOT NULL UNIQUE, 
    "location" TEXT NOT NULL,
    "source" TEXT NOT NULL,
    "frequency" INTEGER NOT NULL,
	PRIMARY KEY (location)
  );`

	statement, err := db.Prepare(createTableSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()

	log.Println("Table created")
}

func InsertLink(name, loc, source string) {
	location, err := GetLink(name, source)

	if location != "" || err != nil {
		log.Fatalln("used name")
		return
	}

	insertLinkSQL := `INSERT INTO navgo(name, location, source, frequency) VALUES (?, ?, ?, ?)`

	statement, err := db.Prepare(insertLinkSQL)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(name, loc, source, 0)

	if err != nil {
		log.Fatalln(err)
	}

}

func GetLink(name, source string) (string, error) {
	row, err := db.Query("SELECT location FROM navgo WHERE name = ? AND source = ?", name, source)

	if err != nil {
		return "", err
	}

	defer row.Close()

	var location string

	for row.Next() {
		err = row.Scan(&location)

		if err != nil {
			log.Fatalln(err)
			return "", err
		}
	}
	return location, nil
}

func DeleteLink(name, source string) bool {
	_, err := db.Exec("DELETE FROM navgo WHERE name = ? AND source = ?", name, source)

	return err == nil
}

func ListLink(source string) ([]string, error) {
	rows, err := db.Query("SELECT name, location FROM navgo WHERE source = ?", source)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var links []string

	for rows.Next() {
		var name string
		var loc string

		err = rows.Scan(&name, &loc)

		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		links = append(links, fmt.Sprintf("%s\t%s\n", name, loc))
	}

	return links, nil

}
