package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Init() {
	connStr := "postgres://todoapp:devpassword@localhost:5432/todos?sslmode=disable"

	var err error
	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateDb()
}

func CreateDb() {
	createTable := `
	CREATE TABLE IF NOT EXISTS todos(
		id SERIAL PRIMARY KEY,
		title TEXT UNIQUE NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false
	)
	`
	_, err := DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func SelectAllTables() {
	query := "SELECT title,done FROM todos"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	type Rows struct {
		title string
		done  bool
	}

	for rows.Next() {
		var data Rows
		err := rows.Scan(&data.title, &data.done)
		if err != nil {
			log.Fatal(err)

		}
		println("Title: "+data.title, "Done: ", data.done)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
func SelectTable(value string) {
	query := "SELECT title,done FROM todos WHERE title =?"
	row := DB.QueryRow(query, &value)

	type Row struct {
		title string
		done  bool
	}

	var data Row
	err := row.Scan(&data.title, &data.done)
	if err != nil {
		log.Fatal(err)

	}
	println("Title: "+data.title, "Done: ", data.done)

}

func InsectTable(value string) {
	query := `INSERT INTO todos(title) VALUES($1)`

	_, err := DB.Exec(query, value)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateTable(keys string, values string) {

	query := `UPDATE todos SET done=$1 WHERE title=$2`

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal("From prepare statement ", err)
	}
	_, err = stmt.Exec(values, keys)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

}
func DeleteTable(values string) {
	_, err := DB.Exec("DELETE FROM todos WHERE title=$1", values)
	if err != nil {
		log.Fatal(err)
	}
}
