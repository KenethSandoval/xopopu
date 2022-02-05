package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type A struct {
	Data string
}

func CSVToJson(f string) (interface{}, error) {
	var (
		dataList []map[string]interface{}
		data     [][]string
		err      error
		database map[string]interface{}
		headers  []string
	)

	data, err = ReadCSV(f)
	if err != nil {
		return nil, err
	}

	headers = append(headers, data[0]...)

	for i, line := range data {
		if i > 0 { // omit header
			database = make(map[string]interface{})

			for j, field := range line {
				database[headers[j]] = field
			}
			dataList = append(dataList, database)
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
