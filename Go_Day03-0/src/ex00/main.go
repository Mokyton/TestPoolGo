package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func main() {
	es7, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(elasticsearch.Version)
	res, err := es7.Info()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println(res)
}
