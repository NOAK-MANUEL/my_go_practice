package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "."
	backupFilePath := "backup.json"
	backupDir := "./backup"

	os.MkdirAll(backupDir, 0755)

	records, err := loadRecords(backupDir)
	if err != nil {
		println(err)
		return
	}

	filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil //avoid files it can't open and directories
		}

		if strings.HasPrefix(path, backupDir) || path == backupFilePath {
			return nil
		}

		checkPermission(path)
		backup(path, records, backupDir)
		return nil
	})
	saveRecords(backupFilePath, records)
}
