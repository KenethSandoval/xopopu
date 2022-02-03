package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func CSVToJson(f string) error {
	err, data := ReadCSV(f)
	if err != nil {
		return err
	}

	m := len(data)
	n := len(data[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v", data[0][0])
		}
	}

	return nil
}

// ReadCSV receive the handler with the name of the last uploaded file
// read the csv line by line and return the read result or an error
func ReadCSV(f string) (error, [][]string) {

	// saves the read of the file in a multidimensional array to be able
	// to extract the header of the csv
	var result [][]string = make([][]string, 0)

	file, err := os.Open("./files/" + f)
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
