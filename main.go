package main

import (
	"flag"
	"log"
)

func main() {
	IsProd := flag.Bool("prod", false, "Sets the environment to production")
	flag.Parse()

	if *IsProd {
		log.Println("Running in production mode")
	} else {
		log.Println("Running in development mode")
	}
}
