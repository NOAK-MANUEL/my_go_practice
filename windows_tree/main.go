package main

import (
	"os"
	"path/filepath"
	"strings"
)

func isParent(index, length int) string {
	if index == length {
		return "└──"
	} else {
		return "├──"
	}
}

func parseTree(path string, depth int) {
	data, err := os.ReadDir(path)
	if err != nil {
		println(err.Error())
		return
	}

	for index, entry := range data {

		if entry.IsDir() {

			println(strings.Repeat("\t", depth), isParent(index, len(data)-1)+entry.Name()+"/")
			parseTree(filepath.Join(path, entry.Name()), depth+1)
		} else {
			println(strings.Repeat("\t", depth), isParent(index, len(data)-1)+entry.Name())

		}
	}
}
func main() {
	path := os.Args[1]

	parseTree(path, 0)
}
