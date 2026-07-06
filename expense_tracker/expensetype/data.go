package expensetype

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Expense struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

func LoadData(filename string) ([]Expense, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, err
	}

	var expenses []Expense

	err = json.Unmarshal(data, &expenses)
	return expenses, err
}

func SaveData(filename string, expense []Expense) error {

	data, err := json.MarshalIndent(expense, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	return os.WriteFile(filename, data, 0644)
}

func ListExpenses(expenses []Expense) {
	for _, data := range expenses {
		fmt.Printf("#%d %-15s %-15s %.2f %s\n", data.ID, data.Name, data.Description, data.Amount, data.Date.Format("2026-12-20"))
	}
}

func Add(expenses []Expense, description string, name string, amount float64) []Expense {
	newExpense := Expense{
		ID:          string(len(expenses) + 1),
		Name:        name,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	return append(expenses, newExpense)
}

func TotalExpenses(expenses []Expense) {
	var total_amount float64 = 0
	for _, data := range expenses {
		total_amount += data.Amount
	}

	fmt.Println(total_amount)
}
func DeleteExpense(id string, filename string, expenses []Expense) {
	newExpenses := []Expense{}
	for _, data := range expenses {
		if data.ID != id {

			newExpenses = append(newExpenses, data)
		}
	}
	SaveData(filename, newExpenses)
}
