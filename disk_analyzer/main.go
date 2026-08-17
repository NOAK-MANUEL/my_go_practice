package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

func getDisk(path string, a, h, r bool, ignore string) (string, int64) {
	if path == "." {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path = dir
	}
	dir_info, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	dirs, err := os.ReadDir(path)
	var totalSize int64 = 0
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range dirs {
		info, err := entry.Info()
		if err == nil {
			among := strings.HasPrefix(info.Name(), ignore)
			if among {
				continue
			}
			if a || r {

				if info.IsDir() {
					if r {
						dirname, size := getDisk(filepath.Join(path, info.Name()), a, h, r, ignore)
						if h {

							println(dir_info.Name()+"/"+dirname, size/1024, "kb")
						} else {

							println(dir_info.Name()+"/"+dirname, size)
						}
						totalSize += size
					} else {

						println(dir_info.Name()+"/"+entry.Name(), "directory")
					}
				} else {
					if h {

						println(dir_info.Name()+"/"+entry.Name(), info.Size()/1024, "kb")
					} else {
						println(dir_info.Name()+"/"+entry.Name(), info.Size())

					}

				}
			}
			totalSize += info.Size()
		}
	}
	return dir_info.Name(), totalSize
}
func main() {
	r := pflag.BoolP("r", "r", false, "recursive search")
	a := pflag.BoolP("a", "a", false, "recursive search")
	h := pflag.BoolP("h", "h", false, "recursive search")
	g := pflag.StringP("ignore", "g", "", "ignore a file e.g .git or .")
	pflag.Parse()
	args := pflag.Args()
	if len(args) < 1 {
		log.Fatal("Incomplete argument")
	}
	path := args[0]
	dirname, size := getDisk(path, *a, *h, *r, *g)
	if *h {
		println(dirname, size/1024, "kb")
	} else {
		println(dirname, size)

	}

}
