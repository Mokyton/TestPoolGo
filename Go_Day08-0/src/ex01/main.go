package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func main() {
	plants := []any{
		UnknownPlant{
			Color:      15,
			LeafType:   "Lancelot",
			FlowerType: "Roses",
		},
		AnotherUnknownPlant{
			FlowerColor: 10,
			LeafType:    "lanceolate",
			Height:      15,
		},
		struct {
			Name   string
			Age    int
			Height int
		}{
			Name:   "checking",
			Age:    123,
			Height: 85,
		},
	}
	describePlant(plants...)
}

func describePlant(plants ...any) {
	for _, plant := range plants {
		t := reflect.TypeOf(plant)
		v := reflect.ValueOf(plant)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Print(field.Name)
			if len(field.Tag) != 0 {
				fmt.Print("(", field.Tag, ")")
			}
			fmt.Println(":", v.FieldByName(field.Name))
		}
	}
}
