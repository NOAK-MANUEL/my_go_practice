package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	return records, err
}
func maxOne(headers []string, rows [][]string, target string) {
	headerIndex := -1
	var previous float64 = 0

	for index, header := range headers {
		if header == target {
			headerIndex = index
		}
	}
	if headerIndex == -1 {
		log.Fatal("Invalid header")
	}

	for _, data := range rows {
		digit, err := strconv.ParseFloat(data[headerIndex], 64)
		if err != nil {
			log.Fatal(err)
		}
		if previous < digit {
			previous = digit
		}
	}

	fmt.Println(target, ":\t", "Max: ", previous)
}
func sumOne(headers []string, rows [][]string, target string) {
	headerIndex := -1
	var previous float64 = 0

	for index, header := range headers {
		if header == target {
			headerIndex = index
		}
	}
	if headerIndex == -1 {
		log.Fatal("Invalid header")
	}

	for _, data := range rows {
		digit, err := strconv.ParseFloat(data[headerIndex], 64)
		if err != nil {
			log.Fatal(err)
		}
		previous += digit

	}

	fmt.Println(target, ":\t", "Sum: ", previous)
}
func main() {

	records, err := readFile("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	headers := records[0]
	data := records[1:]

	sumOne(headers, data, "age")
	maxOne(headers, data, "age")

}
