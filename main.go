package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/convert/csv-to-json", csvToJsonHandler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}

func csvToJsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		http.Error(w, "Failed to parse CSV", http.StatusInternalServerError)
		return
	}

	var jsonData []map[string]string
	headers := records[0]
	for _, row := range records[1:] {
		entry := make(map[string]string)
		for i, value := range row {
			entry[headers[i]] = value
		}
		jsonData = append(jsonData, entry)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}
