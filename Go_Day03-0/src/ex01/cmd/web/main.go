package main

import (
	"ex01/pkg"
	"ex01/pkg/db"
	"fmt"
	"log"
)

func main() {
	test, d, err := db.GetPlaces(10, 10)
	if err != nil {
		log.Fatal(err)
	}

	d := pkg.Store().GetPlaces(10, 20)
	fmt.Println(d)
	fmt.Println(test)
}
