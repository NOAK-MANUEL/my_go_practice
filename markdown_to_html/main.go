package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	var html string = "<html>"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##") {
			content := strings.TrimPrefix(line, "##")
			html = html + "\n" + "<h2>" + content + "</h2>"

		} else if strings.HasPrefix(line, "#") {
			content := strings.TrimPrefix(line, "#")
			html = html + "\n" + "<h1>" + content + "</h1>"

		}

		if strings.HasPrefix(line, ">") {
			content := strings.TrimPrefix(line, ">")
			html = html + "\n" + "<q>" + content + "</q>"

		}
		if strings.HasPrefix(line, "-") {
			content := strings.TrimPrefix(line, "-")
			html = html + "\n" + "<li>" + content + "</li>"

		}

		if strings.Contains(line, "**") {
			lineSlice := strings.Fields(line)

			for index, text := range lineSlice {
				if strings.HasPrefix(text, "**") {
					content := strings.TrimPrefix(text, "**")
					content = strings.TrimSuffix(content, "**")
					lineSlice[index] = "<b>" + content + "</b>"
				}
			}
			html = html + "\n" + strings.Join(lineSlice, " ")
		}

		// lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	file.Close()
	file, err = os.Create("index.html")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	html = html + "\n" + "</html>"
	file.WriteString(html)

}
