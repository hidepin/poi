package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	var (
		noexec bool
	)
	flag.BoolVar(&noexec, "n", false, "no exec")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		fatalErr = err
		return
	}

	err = filepath.Walk(dir, clean_up)

	if err != nil {
		fatalErr = err
		return
	}

	fmt.Println("poi.")
}

func clean_up(path string, fi os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
