package main

import "strings"

func render(template string, data map[string]string) string {
	var result strings.Builder
	var i = 0
	for i < len(template) {
		if i+1 < len(template) && template[i] == '{' && template[i+1] == '{' {
			//found starting point of }
			end := strings.Index(template[i:], "}}")
			if end == -1 {
				result.WriteString(template[i:])
				break
			}
			key := strings.TrimSpace(template[i+2 : i+end])
			value, exist := data[key]
			if exist {
				result.WriteString(value)
			} else {
				result.WriteString("{{")
				result.WriteString(key)
				result.WriteString("}}")
			}

			i += end + 2
		} else {
			result.WriteByte(template[i])
			i++
		}

	}
	return result.String()

}
