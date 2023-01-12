package main

import (
	"ex01/compare"
	"flag"
	"log"
)

func main() {
	fOld := flag.String("old", "original_database.xml", "path to original_database")
	fNew := flag.String("new", "stolen_database.json", "path to stolen_database")
	flag.Parse()
	err := compare.Compare(*fOld, *fNew)
	if err != nil {
		log.Fatal(err)
	}

}
