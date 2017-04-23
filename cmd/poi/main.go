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
		help = *flag.Bool("h", false, "help")
	)
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}
	fmt.Println("poi.")
}
