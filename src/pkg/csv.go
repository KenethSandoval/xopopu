package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// TestCSV struct the test for csv to json
type TestCSV struct {
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int    `json:"rank"`
}

// TODO: working...
func CSVToJson(f string) ([]TestCSV, error) {
	var (
		dataList []TestCSV
		data     [][]string
		err      error
	)

	data, err = ReadCSV(f)
	if err != nil {
		return nil, err
	}

	for i, line := range data {
		if i > 0 { // omit header
			var rec TestCSV
			for j, field := range line {
				switch j {
				case 0:
					rec.Vegetable = field
				case 1:
					rec.Fruit = field
				case 2:
					rec.Rank, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				}
			}
			dataList = append(dataList, rec)
		}
	}

	return dataList, nil
}

// ReadCSV receive the handler with the name of the last uploaded file
// read the csv line by line and return the read result or an error
func ReadCSV(f string) ([][]string, error) {
	// saves the read of the file in a multidimensional array to be able
	// to extract the header of the csv
	result := make([][]string, 0)

	file, err := os.Open("./files/" + f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		result = append(result, row)
	}
	return result, nil
}
