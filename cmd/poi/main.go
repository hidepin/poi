package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hidepin/poi"
)

var (
	noexec   bool
	root_dir string
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
		os.Exit(0)
	}()
	dir, err := os.Getwd()
	if err != nil {
		fatalErr = err
		return
	}

	flag.StringVar(&root_dir, "d", dir, "exec dir")
	flag.BoolVar(&noexec, "n", false, "no exec")
	flag.Parse()

	err = filepath.Walk(root_dir, clean_up)

	if err != nil {
		fatalErr = err
		return
	}

	fmt.Println("poi done.")
}

func clean_up(path string, fi os.FileInfo, err error) error {
	var cleaner poi.Cleaner
	if noexec {
		cleaner = poi.NOEXEC
	} else {
		cleaner = poi.EXEC
	}

	if fi.IsDir() && path != root_dir {
		return filepath.SkipDir
	} else if !fi.IsDir() && strings.HasSuffix(fi.Name(), "~") {
		cleaner.Clean(path)
	}
	return nil
}
