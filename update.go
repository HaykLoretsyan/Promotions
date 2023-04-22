package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

func UpdateData(csvFile string) error {

	data, err := readCsvFile(csvFile)
	if err != nil {
		return err
	}

	models, err := modelsFromData(data)
	if err != nil {
		return err
	}

	if err = RemoveAll(); err != nil {
		return err
	}
	return SaveList(models)
}

func readCsvFile(filePath string) ([][]string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("unable to read input file " + filePath + err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New("unable to parse file as CSV for " + filePath + err.Error())
	}

	return records, nil
}

func modelsFromData(data [][]string) ([]Model, error) {

	models := make([]Model, len(data))
	for i, entry := range data {
		if len(entry) != 3 {
			return nil, errors.New("illegal csv file: wrong number of fields at line " + strconv.Itoa(i+1))
		}
		models[i].Id = entry[0]
		value, err := strconv.ParseFloat(entry[1], 32)
		if err != nil {
			return nil, errors.New("illegal csv file: " + err.Error())
		}
		models[i].Price = float32(value)
		models[i].ExpirationDate = entry[2]
	}
	return models, nil
}
