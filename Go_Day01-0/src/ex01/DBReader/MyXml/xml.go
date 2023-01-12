package MyXml

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    Cake     `xml:"cake"`
}

type Cake []struct {
	XMLName     xml.Name    `xml:"cake" json:"-"`
	Name        string      `xml:"name"`
	Stovetime   string      `xml:"stovetime"`
	Ingredients Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	Text xml.Name `xml:"ingredients" json:"-"`
	Item []Item   `xml:"item"`
}

type Item struct {
	XMLName   xml.Name `xml:"item" json:"-"`
	Itemname  string   `xml:"itemname"`
	Itemcount string   `xml:"itemcount"`
	Itemunit  string   `xml:"itemunit"`
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
	err = xml.Unmarshal(byteValue, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Recipes) Convert() ([]byte, error) {
	byteValue, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

func (c *Recipes) CreateAnotherExt(data []byte) error {
	file, err := os.Create("fromXmlTo.json")
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
