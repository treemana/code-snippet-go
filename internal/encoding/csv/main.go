package main

import (
	"encoding/csv"
	"os"
)

func writeCSV(path string, data [][]string) error {
	if len(path) == 0 || data == nil {
		return nil
	}

	csvFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = csvFile.Close()
	}()

	writer := csv.NewWriter(csvFile)
	for _, line := range data {
		if err = writer.Write(line); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func main() {
	_ = writeCSV("", nil)
}
