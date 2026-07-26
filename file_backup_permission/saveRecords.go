package main

import (
	"encoding/json"
	"os"
)

func saveRecords(path string, records BackupRecords) {
	data, err := json.MarshalIndent(records, "", "\t")
	if err != nil {
		println("Couldn't parsed record: ", err)
		return
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		println("Couldn't save record: ", err)
		return
	}

}
