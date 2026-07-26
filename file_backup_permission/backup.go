package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func backup(path string, records BackupRecords, backupDir string) {
	newHash, err := hash(path)
	if err != nil {
		println(err)
		return
	}

	oldHash, exist := records[path]
	if exist || oldHash == newHash {
		println("File the same")
		return
	}

	err = copyFile(path, backupDir)
	if err != nil {
		println("Error at line 24: ", err)
		return
	}

	fmt.Printf("Backed up: %s (%s)\n", path, reason(exist))
	records[path] = newHash

}

func reason(existed bool) string {
	if existed {

		return "changed"
	}
	return "new"

}

func copyFile(path, backupDir string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	dist := filepath.Join(backupDir, filepath.Base(path))

	return os.WriteFile(dist, data, 0644)
}
