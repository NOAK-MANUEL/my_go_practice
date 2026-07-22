package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	config := make(Config)
	var head string
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" || strings.HasPrefix(text, "#") || strings.HasPrefix(text, ";") {
			//ignoring comment
			continue
		}

		if strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") {
			head = text[1 : len(text)-1]
			config[head] = make(map[string]string)
			continue
		}

		if strings.Contains(text, "=") {
			if head == "" {
				continue
			}
			value := strings.SplitN(text, "=", 2) //gives only 2 length of a slice
			valHead := strings.TrimSpace(value[0])
			valBody := strings.TrimSpace(value[1])

			config[head][valHead] = valBody
		}

	}

	return config, scanner.Err()
}
