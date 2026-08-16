package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func find[T comparable](slice []T, target T) (int, bool) {
	for index, content := range slice {
		if content == target {
			return index, true
		}
	}
	return -1, false
}

func main() {

	case_insensitive := pflag.Bool("i", false, "case insensitive")
	show_line := pflag.Bool("n", false, "show index")
	pflag.Parse()
	args := pflag.Args()
	if len(args) < 2 {
		log.Fatal("Expected 2 command <target> <source>")
	}

	target := args[0]
	source := args[1]
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if *case_insensitive {
		target = strings.ToLower(target)
	}
	var lineIndex = 1
	for scanner.Scan() {
		line := scanner.Text()
		searchLine := line

		if *case_insensitive {
			searchLine = strings.ToLower(searchLine)
		}

		if strings.Contains(searchLine, target) {
			if *show_line {

				println(lineIndex, line)
			} else {
				println(lineIndex, line)

			}
		}
		lineIndex++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
