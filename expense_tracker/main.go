package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/NOAK-MANUEL/my_go_practice/expense_tracker/expensetype"
)

func main() {
	filename := "data.json"
	expenses, err := expensetype.LoadData(filename)
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Argument can be less than 2")
	}

	switch args[2] {
	case "add":
		if len(args) < 5 {
			log.Fatal("Usage: expense add <description> <amount>")
		}
		amount, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			log.Fatal(err)
		}
		description := args[3]
		fmt.Println(description)
		expenses = expensetype.Add(expenses, description, args[1], amount)
		err = expensetype.SaveData(filename, expenses)
		if err != nil {
			log.Fatal(err)
		}
	case "list":
		expensetype.ListExpenses(expenses)
	case "total":
		expensetype.TotalExpenses(expenses)
	case "delete":
		expensetype.DeleteExpense(args[1], filename, expenses)
		fmt.Print("Deleted successfully")

	}
}
