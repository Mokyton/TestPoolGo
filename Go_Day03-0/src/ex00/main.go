package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/grailbio/base/tsv"
	"github.com/olivere/elastic"
	"log"
	"os"
	"strconv"
	"strings"
)

type Doc struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

func main() {
	pathToDataSets := flag.String("dSet", "../../materials/data.csv", "path to csv file")
	indexName := flag.String("iName", "places", "index name")
	countOfData := flag.Int("cData", 0, "count of data u want to upload")

	flag.Parse()
	dataSet, err := getDataSet(*pathToDataSets)
	if *countOfData == 0 {
		*countOfData = len(dataSet)
	}
	if err != nil {
		log.Fatal(err)
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < *countOfData; i++ {

		place := setData(dataSet[i])
		myJson, err := jsonHandler(place)
		if err != nil {
			log.Fatal(err)
		}

		request := esapi.IndexRequest{
			Index:        *indexName,
			DocumentID:   strconv.Itoa(i + 1),
			DocumentType: "place",
			Body:         strings.NewReader(string(myJson)),
			Refresh:      "true",
		}
		response, err := request.Do(context.Background(), es)
		if err != nil {
			log.Fatal(err)
		}

		if response.IsError() {
			log.Fatalln("Error indexing document")
		}

		fmt.Println("status:", response.Status())
		response.Body.Close()
	}

}

func getDataSet(fileCSV string) ([][]string, error) {
	file, err := os.Open(fileCSV)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var d []byte
	reader := tsv.NewReader(file)
	reader.Read(d)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	data = data[1:][:]
	return data, nil
}

func jsonHandler(place Doc) ([]byte, error) {
	res, err := json.Marshal(&place)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func setData(data []string) Doc {
	id, _ := strconv.Atoi(data[0])
	lon, _ := strconv.ParseFloat(data[4], 64)
	lat, _ := strconv.ParseFloat(data[5], 64)
	return Doc{
		ID:       id,
		Name:     data[1],
		Address:  data[2],
		Phone:    data[3],
		Location: elastic.GeoPoint{Lon: lon, Lat: lat},
	}
}
