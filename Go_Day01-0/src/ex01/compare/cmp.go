package compare

import (
	"errors"
	"ex01/DBReader/MyJson"
	"ex01/DBReader/MyXml"
	"fmt"
	"strings"
)

type recipes struct {
	Cake []cake
}

type cake struct {
	Name        string
	Time        string
	Ingredients []ingredients
}

type ingredients struct {
	IngredientName  string
	IngredientCount string
	IngredientUnit  string
}

var (
	newCakeNames []string
	oldCakeNames []string
)

func handleXml(fileName string) (*recipes, error) {
	local := recipes{}
	xml, err := MyXml.Parse(fileName)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(xml.Cake); i++ {
		Cake := cake{xml.Cake[i].Name, xml.Cake[i].Stovetime, nil}
		for j := 0; j < len(xml.Cake[i].Ingredients.Item); j++ {
			Ingredients := ingredients{xml.Cake[i].Ingredients.Item[j].Itemname, xml.Cake[i].Ingredients.Item[j].Itemcount, xml.Cake[i].Ingredients.Item[j].Itemunit}
			Cake.Ingredients = append(Cake.Ingredients, Ingredients)
		}
		local.Cake = append(local.Cake, Cake)
	}
	return &local, nil
}

func handleJson(fileName string) (*recipes, error) {
	local := recipes{}
	json, err := MyJson.Parse(fileName)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(json.Cake); i++ {
		Cake := cake{
			json.Cake[i].Name,
			json.Cake[i].Time,
			nil,
		}
		for j := 0; j < len(json.Cake[i].Ingredients); j++ {
			Ingredients := ingredients{
				json.Cake[i].Ingredients[j].IngredientName,
				json.Cake[i].Ingredients[j].IngredientCount,
				json.Cake[i].Ingredients[j].IngredientUnit,
			}
			Cake.Ingredients = append(Cake.Ingredients, Ingredients)
		}
		local.Cake = append(local.Cake, Cake)
	}
	return &local, err
}

func Compare(origName, stolenName string) error {
	oldCake, err := checkSuffix(origName)
	if err != nil {
		return err
	}
	newCake, err := checkSuffix(stolenName)
	if err != nil {
		return err
	}
	for i := 0; i < len(oldCake.Cake); i++ {
		oldCakeNames = append(oldCakeNames, oldCake.Cake[i].Name)
	}
	for i := 0; i < len(newCake.Cake); i++ {
		newCakeNames = append(newCakeNames, newCake.Cake[i].Name)
	}
	cakeAdded()
	cakeRemoved()
	cakesCompare(oldCake, newCake)
	return nil
}

func checkSuffix(fileName string) (*recipes, error) {
	if strings.HasSuffix(fileName, ".xml") {
		return handleXml(fileName)
	} else if strings.HasSuffix(fileName, ".json") {
		return handleJson(fileName)
	}
	return nil, errors.New("Wrong file extension ")
}

func cakeAdded() {
	for i := 0; i < len(newCakeNames); i++ {
		if !arrContain(newCakeNames[i], oldCakeNames) {
			fmt.Printf("ADDED cake \"%s\"\n", newCakeNames[i])
		}
	}
}

func cakeRemoved() {
	for i := 0; i < len(oldCakeNames); i++ {
		if !arrContain(oldCakeNames[i], newCakeNames) {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCakeNames[i])
		}
	}
}

func arrContain(cake string, src []string) bool {
	for _, v := range src {
		if cake == v {
			return true
		}
	}
	return false
}

func cakesCompare(oldCake, newCake *recipes) {
	oldMap := make(map[string]cake, len(oldCake.Cake))
	newMap := make(map[string]cake, len(newCake.Cake))
	for i := 0; i < len(oldCakeNames); i++ {
		oldMap[oldCake.Cake[i].Name] = oldCake.Cake[i]
	}
	for i := 0; i < len(newCakeNames); i++ {
		newMap[newCake.Cake[i].Name] = newCake.Cake[i]
	}
	for i := 0; i < len(oldCakeNames); i++ {
		oldV := oldMap[oldCakeNames[i]]
		newV := newMap[oldCakeNames[i]]
		if oldV.Name == newV.Name {
			oldIngrMap := make(map[string]ingredients, len(oldV.Ingredients))
			newIngrMap := make(map[string]ingredients, len(newV.Ingredients))
			if oldV.Time != newV.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					oldV.Name, newV.Time, oldV.Time)
			}
			ingredientsAdded(oldV, newV)
			ingredientsRemoved(oldV, newV)
			for j := 0; j < len(oldV.Ingredients); j++ {
				oldIngrMap[oldV.Ingredients[j].IngredientName] = oldV.Ingredients[j]
			}
			for j := 0; j < len(newV.Ingredients); j++ {
				newIngrMap[newV.Ingredients[j].IngredientName] = newV.Ingredients[j]
			}
			oIV := oldIngrMap[oldV.Ingredients[i].IngredientName]
			nIV := newIngrMap[newV.Ingredients[i].IngredientName]
			if oIV.IngredientUnit != nIV.IngredientUnit {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
					oIV.IngredientName, oldV.Name, nIV.IngredientUnit, oIV.IngredientUnit)
			}
			if oIV.IngredientCount != nIV.IngredientCount {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
					oIV.IngredientName, oldV.Name, nIV.IngredientCount, oIV.IngredientCount)
			}
		}
	}
}

func ingredientsAdded(oldIngr, newIngr cake) {
	oldIngrNames := make([]string, 0, len(newIngr.Ingredients))
	for i := 0; i < len(oldIngr.Ingredients); i++ {
		oldIngrNames = append(oldIngrNames, oldIngr.Ingredients[i].IngredientName)
	}
	for i := 0; i < len(newIngr.Ingredients); i++ {
		if !arrContain(newIngr.Ingredients[i].IngredientName, oldIngrNames) {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n",
				newIngr.Ingredients[i].IngredientName, newIngr.Name)
		}
	}
}

func ingredientsRemoved(oldIngr, newIngr cake) {
	newIngrNames := make([]string, 0, len(newIngr.Ingredients))
	for i := 0; i < len(newIngr.Ingredients); i++ {
		newIngrNames = append(newIngrNames, newIngr.Ingredients[i].IngredientName)
	}
	for i := 0; i < len(oldIngr.Ingredients); i++ {
		if !arrContain(oldIngr.Ingredients[i].IngredientName, newIngrNames) {
			fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", oldIngr.Ingredients[i].IngredientName, oldIngr.Name)
		}
	}
}
