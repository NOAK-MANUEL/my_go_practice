package main

import (
	"fmt"
	"os"
)

func checkPermission(path string) {
	info, err := os.Stat(path)
	if err != nil {
		println("Error checking permission: ", err)
		return
	}
	perm := info.Mode().Perm()
	if perm&0002 != 0 {
		fmt.Printf("%s has world-writable permissions %o", info.Name(), perm)
	}
}
