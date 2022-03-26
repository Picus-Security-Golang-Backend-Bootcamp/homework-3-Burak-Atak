package helper

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Csv file could not open.")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvSlice, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Csv file could not read.")
	}

	return csvSlice
}
