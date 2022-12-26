package main

import (
	"ex01/compare"
	"flag"
)

func main() {
	fOld := flag.String("old", "original_database.xml", "path to original_database")
	fNew := flag.String("new", "stolen_database.json", "path to stolen_database")
	flag.Parse()
	//f1, _ := MyXml.Parse(*fOld)
	//f2, _ := MyJson.Parse(*fNew)

	compare.Compare(*fOld, *fNew)

}
