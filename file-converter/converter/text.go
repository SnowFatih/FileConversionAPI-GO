package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func CSVToJSON(inputPath string) (string, error) {
	outputPath := inputPath[:len(inputPath)-4] + ".json"

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	data := []map[string]string{}
	headers := records[0]

	for _, row := range records[1:] {
		entry := make(map[string]string)
		for i, value := range row {
			entry[headers[i]] = value
		}
		data = append(data, entry)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return "", err
	}

	fmt.Println("✅ Conversion CSV → JSON réussie :", outputPath)
	return outputPath, nil
}

func JSONToCSV(inputPath string) (string, error) {
	outputPath := inputPath[:len(inputPath)-5] + ".csv"

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var data []map[string]string
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return "", err
	}

	headers := []string{}
	for key := range data[0] {
		headers = append(headers, key)
	}

	csvData := [][]string{headers}
	for _, entry := range data {
		row := []string{}
		for _, header := range headers {
			row = append(row, entry[header])
		}
		csvData = append(csvData, row)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	err = writer.WriteAll(csvData)
	if err != nil {
		return "", err
	}
	writer.Flush()

	fmt.Println("✅ Conversion JSON → CSV réussie :", outputPath)
	return outputPath, nil
}
