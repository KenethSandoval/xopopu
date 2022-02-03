package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// ReadCSV receive the handler with the name of the last uploaded file
// read the csv line by line and return the read result or an error
func ReadCSV() (error, [][]string) {
	var result [][]string
	file, err := os.Open("./files/test.csv")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			return err, nil
		}
		if err == io.EOF {
			break
		}
		result = append(result, row)
	}
	return nil, result
}
