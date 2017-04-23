package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hidepin/poi"
)

var (
	noexec    bool
	recursive bool
	root_dir  string
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

	flag.BoolVar(&noexec, "n", false, "no exec")
	flag.BoolVar(&recursive, "r", false, "recursive")
	flag.StringVar(&root_dir, "d", dir, "exec dir")
	flag.Parse()
	if flag.NArg() != 0 {
		fatalErr = errors.New("arguments error")
		return
	}

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

	if fi.IsDir() {
		if !recursive && path != root_dir {
			return filepath.SkipDir
		}
	} else if strings.HasSuffix(fi.Name(), "~") {
		cleaner.Clean(path)
	}
	return nil
}
