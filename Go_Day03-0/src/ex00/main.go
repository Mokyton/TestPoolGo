package main

import (
	"context"
	"encoding/json"
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
	Name     string           `json:"name""`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

//var (
//	indexName string
//	//numWorkers int
//	//flushBytes int
//	//numItems   int
//)

func main() {
	dataSet, err := getDataSet("../../materials/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	//bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
	//	Index:  indexName,
	//	Client: es,
	//	//NumWorkers: numWorkers,
	//	//FlushBytes: int(flushBytes),
	//	//FlushInterval: 30 * time.Second,
	//})

	for i := 0; i < 25; i++ {
		place := setData(dataSet[i])
		myJson, err := jsonHandler(place)
		if err != nil {
			log.Fatal(err)
		}

		request := esapi.IndexRequest{
			Index:      "places",
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(string(myJson)),
			Refresh:    "true",
		}
		response, err := request.Do(context.Background(), es)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		if response.IsError() {
			log.Fatalln("Error indexing document")
		}

		var res map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
			log.Fatalf("error parsing the response body: %s", err)
		}

		fmt.Println("status:", response.Status())
	}
	//
	//if err = bi.Close(context.Background()); err != nil {
	//	log.Fatalf("Unexpected error: %s", err)
	//}

}

func getDataSet(fileCSV string) ([][]string, error) {
	file, err := os.Open(fileCSV)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := tsv.NewReader(file)
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
