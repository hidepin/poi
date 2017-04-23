package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	f, err := os.Open(dir)
	if err != nil {
		fatalErr = err
		return
	}
	defer f.Close()
	fis, err := f.Readdir(0)
	for _, fi := range fis {
		if !fi.IsDir() {
			fmt.Println(fi.Name())
		}
	}

	fmt.Println("poi.")
}
