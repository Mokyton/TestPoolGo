package main

import (
	"DAY01.com/DBReader"
	"DAY01.com/DBReader/MyJson"
	"DAY01.com/DBReader/MyXml"
	"errors"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	fName := flag.String("f", "", "path to file with json/xml extension")
	flag.Parse()
	file, err := chooseExtension(*fName)
	if err != nil {
		log.Fatal(err)
	}
	err = start(file)
	if err != nil {
		log.Fatal(err)
	}
}

func chooseExtension(fileName string) (DBReader.DBReader, error) {
	var res DBReader.DBReader
	if strings.HasSuffix(fileName, ".json") {
		res = MyJson.New()
	} else if strings.HasSuffix(fileName, ".xml") {
		res = MyXml.New()
	} else {
		return nil, errors.New("wrong extension choose another file")
	}
	fp, err := os.Open(fileName)
	defer fp.Close()
	if err != nil {
		return nil, err
	}
	err = res.Parse(fp)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func start(reader DBReader.DBReader) error {
	data, err := reader.Convert()
	if err != nil {
		return err
	}
	err = reader.CreateAnotherExt(data)
	if err != nil {
		return err
	}
	return nil
}
