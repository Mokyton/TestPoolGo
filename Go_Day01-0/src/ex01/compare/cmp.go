package compare

import (
	"ex01/DBReader/MyJson"
	"ex01/DBReader/MyXml"
	"fmt"
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
	//for k := 0; k < len(local.Cake); k++ {
	//	sort.Slice(local.Cake[k].Ingredients, func(i, j int) bool {
	//		return local.Cake[k].Ingredients[i].IngredientName < local.Cake[k].Ingredients[j].IngredientName
	//	})
	//}
	//sort.Slice(local.Cake, func(i, j int) bool {
	//	return local.Cake[i].Name < local.Cake[j].Name
	//})
	return &local, nil
}

func handleJson(fileName string) (*recipes, error) {
	local := recipes{}
	json, err := MyJson.Parse(fileName)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(json.Cake); i++ {
		Cake := cake{json.Cake[i].Name, json.Cake[i].Time, nil}
		for j := 0; j < len(json.Cake[i].Ingredients); j++ {
			Ingredients := ingredients{json.Cake[i].Ingredients[j].IngredientName, json.Cake[i].Ingredients[j].IngredientCount, json.Cake[i].Ingredients[j].IngredientUnit}
			Cake.Ingredients = append(Cake.Ingredients, Ingredients)
		}
		local.Cake = append(local.Cake, Cake)
	}
	//for k := 0; k < len(local.Cake); k++ {
	//	sort.Slice(local.Cake[k].Ingredients, func(i, j int) bool {
	//		return local.Cake[k].Ingredients[i].IngredientName < local.Cake[k].Ingredients[j].IngredientName
	//	})
	//}
	//sort.Slice(local.Cake, func(i, j int) bool {
	//	return local.Cake[i].Name < local.Cake[j].Name
	//})
	return &local, err
}

func Compare(firstName, secondName string) error {
	firstFile, _ := handleJson(secondName)
	secondFile, _ := handleXml(firstName)
	cakeAdd(firstFile, secondFile)
	cakeRemoved(firstFile, secondFile)
	return nil
}

func cakeRemoved(oldCake, newCake *recipes) {
	var oldCakeNames []string
	for i := 0; i < len(oldCake.Cake); i++ {
		oldCakeNames = append(oldCakeNames, oldCake.Cake[i].Name)
	}
	for i := 0; i < len(newCake.Cake); i++ {
		if !arrContain(newCake.Cake[i].Name, oldCakeNames) {
			fmt.Printf("REMOVED cake \"%s\"\n", newCake.Cake[i].Name)
		}
	}
}

func cakeAdd(oldCake, newCake *recipes) {
	var newCakeNames []string
	for i := 0; i < len(newCake.Cake); i++ {
		newCakeNames = append(newCakeNames, newCake.Cake[i].Name)
	}
	for i := 0; i < len(oldCake.Cake); i++ {
		if !arrContain(oldCake.Cake[i].Name, newCakeNames) {
			fmt.Printf("ADDED cake \"%s\"\n", oldCake.Cake[i].Name)
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
