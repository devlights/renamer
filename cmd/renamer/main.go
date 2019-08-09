package main

import (
	"flag"
	"github.com/devlights/renamer/lib"
	"log"
	"os"
	"path/filepath"
)

var (
	baseDir string
	verbose bool
)

func logf(format string, v ...interface{}) {
	if verbose {
		log.Printf(format, v...)
	}
}

func main() {

	var (
		oldPattern string
		newPattern string
		recursive  bool
	)

	flag.StringVar(&baseDir, "path", ".", "Target directory")
	flag.StringVar(&oldPattern, "old", "", "Old pattern")
	flag.StringVar(&newPattern, "new", "", "New pattern")
	flag.BoolVar(&recursive, "r", false, "Perform recursively")
	flag.BoolVar(&verbose, "v", false, "Output verbose log messages")

	flag.Parse()

	if recursive {
		err := filepath.Walk(baseDir, func(p string, info os.FileInfo, e error) error {
			if e != nil {
				log.Fatal(e)
			}

			if info.IsDir() {
				if verbose {
					log.Printf("Enter [%s]", p)
				}
				if renameErr := lib.Rename(p, oldPattern, newPattern, verbose); renameErr != nil {
					return renameErr
				}
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	} else {
		if err := lib.Rename(baseDir, oldPattern, newPattern, verbose); err != nil {
			log.Fatal(err)
		}
	}

	os.Exit(0)
}
