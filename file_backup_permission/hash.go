package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func hash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
