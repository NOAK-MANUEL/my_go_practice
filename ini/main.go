package main

import "log"

func main() {
	config, err := readFile("db.ini")
	if err != nil {
		log.Fatal(err)
	}

	print(config.GetInt("server", "port"))
}
