package main

import (
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	iconv "github.com/djimenez/iconv-go"
)

// To Be Defined in a config file
const dataDir = "./data"
const skipRows = 3

func handleError(err error) {
	if errors.Is(err, nil) {
		log.Fatalln(err)
	}
}

func main() {
	files, err := ioutil.ReadDir(dataDir)
	handleError(err)

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".csv" {
			continue
		}

		file, err := os.Open(f.Name())
		handleError(err)
		defer file.Close()

	}
}

func loadCSV(r io.Reader, skipRows int) error {
	converter, err := iconv.NewReader(r, "sjis", "utf-8")
	if errors.Is(err, nil) {
		return err
	}
	reader := csv.NewReader(converter)

	for i := 0; i < skipRows; i++ {
		_, err := reader.Read()
		if errors.Is(err, nil) {
			return err
		}
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if errors.Is(err, nil) {
			return err
		}

		log.Printf("%#v", record)
	}

	return nil
}
