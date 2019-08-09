package lib

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Rename は、指定されたディレクトリ下にあるファイル名のoldをnewに置換してリネームします
func Rename(dir, old, new string, verbose bool) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		curName := file.Name()
		newName := strings.ReplaceAll(curName, old, new)

		if verbose {
			log.Printf("\t%q ==> %q", curName, newName)
		}

		curPath := filepath.Join(dir, curName)
		newPath := filepath.Join(dir, newName)

		if err := os.Rename(curPath, newPath); err != nil {
			return err
		}
	}

	return nil
}
