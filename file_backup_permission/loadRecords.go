package main

import (
	"encoding/json"
	"os"
)

type BackupRecords map[string]string

func loadRecords(path string) (BackupRecords, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return make(BackupRecords), nil
	}

	records := make(BackupRecords)
	err = json.Unmarshal(data, &records)
	return records, err
}
