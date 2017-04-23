package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()
	var (
		list bool
	)
	flag.BoolVar(&noexec, "n", false, "no exec")
	flag.Parse()

	fmt.Println("poi.")
}
