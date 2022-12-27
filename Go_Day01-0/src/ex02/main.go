package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	oldFile := flag.String("old", "txtFiles/snapshot1.txt", "path to old database")
	newFile := flag.String("new", "txtFiles/snapshot2.txt", "path to new database")
	flag.Parse()
	newF, oldF, err := OpenFiles(*oldFile, *newFile)
	if err != nil {
		log.Fatal(err)
	}
	err = CompareFiles(newF, oldF)
	if err != nil {
		log.Fatal(err)
	}
}

func OpenFiles(firstFile, secondFile string) (*os.File, *os.File, error) {
	firFile, err := os.Open(firstFile)
	if err != nil {
		return nil, nil, err
	}
	secFile, err := os.Open(secondFile)
	if err != nil {
		return nil, nil, err
	}
	return firFile, secFile, nil
}

func ReadFile(file *os.File) ([]string, error) {
	var res []string
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		res = append(res, scan.Text())
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	file.Seek(0, 0)
	return res, nil
}

func CompareFiles(firstFile, secondFile *os.File) error {
	defer firstFile.Close()
	defer secondFile.Close()
	var secondFileData []string
	firstFileData, err := ReadFile(firstFile)
	sort.Strings(firstFileData)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(secondFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			secondFileData = append(secondFileData, line)
			if !stringContains(line, firstFileData) {
				fmt.Printf("ADDED %s\n", line)
			}
		}

	}
	if err = scanner.Err(); err != nil {
		return err
	}
	sort.Strings(secondFileData)
	for i := 0; i < len(firstFileData); i++ {
		if firstFileData[i] != " " {
			if !stringContains(firstFileData[i], secondFileData) {
				fmt.Printf("REMOVED %s\n", firstFileData[i])
			}
		}
	}

	return nil
}

func stringContains(elem string, data []string) bool {

	i := sort.SearchStrings(data, elem)
	if i < len(data) && data[i] == elem {
		return true
	}

	return false
}
