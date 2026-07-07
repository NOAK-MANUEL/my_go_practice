package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	store := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("e.g Set/Delete/Get  name Manuel")
	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)

		lineSlice := strings.Fields(line)
		fmt.Println(lineSlice)

		name := lineSlice[1]
		command := lineSlice[0]
		switch strings.ToLower(command) {
		case "set":
			data := lineSlice[2]
			store[name] = data
			fmt.Println(">", "OK")
		case "get":
			validData, ok := store[name]
			if !ok {
				fmt.Println(">", "Input not found")
			} else {

				fmt.Println(">", validData)
			}
		case "delete":
			delete(store, name)
			fmt.Println(">", "OK")
		default:
			fmt.Println("> Invalid input")

		}

		fmt.Print("> ")

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
