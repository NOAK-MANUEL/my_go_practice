package main

import (
	"log"
	"strconv"
)

type Config map[string]map[string]string

func (c Config) GetInt(section, value string) int {
	num, err := strconv.Atoi(c[section][value])
	if err != nil {
		log.Fatal(err)
	}
	return num
}
func (c Config) GetBool(section, value string) bool {
	boo, err := strconv.ParseBool(c[section][value])
	if err != nil {
		log.Fatal(err)
	}
	return boo
}
