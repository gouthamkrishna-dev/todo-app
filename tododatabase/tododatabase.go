package tododatabase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"fmt"
)

var DB *sql.DB

func Createdatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "todo.db")
	if err != nil {
		fmt.Printf("database not connected\n")
	}
	pingerr := DB.Ping()

	if pingerr != nil {
		fmt.Printf("database not connected while ping")
	}
	fmt.Printf("Successfully connected")
	Adddatabase()
}

func Adddatabase() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS todo (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description text,
			status text Not Null,
			priority text Not Null,
			created_at DATETIME NOT NULL DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime'))
		);`)

	if err != nil {
		fmt.Printf("Message 1\n %v", err)
	}

}
