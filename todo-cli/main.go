package main

import (
	"log"
	"os"
	"todo-cli/db"
)

func main() {

	db.Init()
	value := os.Args[2]
	switch os.Args[1] {
	case "add":
		db.InsectTable(value)
		println("Created")
	case "delete":
		db.DeleteTable(value)
		println("Row Deleted")
	case "edit":
		if len(os.Args) < 3 {
			log.Fatal("Expected 3 argument, got 2")
		}
		db.UpdateTable(os.Args[2], value)
		println("Row Edited")
	case "select":
		if value == "*" {
			db.SelectAllTables()
		} else {

			db.SelectTable(value)
		}
	}

}
