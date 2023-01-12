package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	path := flag.String("a", "./", "place where archive should be created")
	flag.Parse()
	files := os.Args[1:]
	var wg sync.WaitGroup
	if len(files) == 0 {
		log.Fatalln("error add some args")
	}
	for _, v := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := createArchive(v, *path)
			if err != nil {
				log.Fatal(err)
			}
		}()

	}
	wg.Wait()

}

func createArchive(filename, path string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	buf, err := os.Create(fmt.Sprintf("%s/%s_%d.tar.gz",
		path, fileNameWithoutExtSliceNotation(filename), info.ModTime().Unix()))
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing archive: %v", err))
	}
	defer buf.Close()

	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	header.Name = filename
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
