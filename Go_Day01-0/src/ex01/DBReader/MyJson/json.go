package MyJson

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Recipes struct {
	Cake []Cake `json:"cake"`
}

type Cake struct {
	Name        string        `json:"name"`
	Time        string        `json:"time"`
	Ingredients []Ingredients `json:"ingredients"`
}

type Ingredients struct {
	IngredientName  string `json:"ingredient_name"`
	IngredientCount string `json:"ingredient_count"`
	IngredientUnit  string `json:"ingredient_unit,omitempty"`
}

func New() *Recipes {
	return &Recipes{}
}

func Parse(fileName string) (*Recipes, error) {
	res := New()
	f, err := os.Open(fileName)
	defer f.Close()
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Recipes) Convert() ([]byte, error) {
	byteValue, err := xml.MarshalIndent(c, "", "    ")
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

func (c *Recipes) CreateAnotherExt(data []byte) error {
	file, err := os.Create("fromJsonToXml.xml")
	defer file.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
