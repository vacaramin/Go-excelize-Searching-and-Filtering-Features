package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuri/excelize/v2"
)

type ExcelRow struct {
	Cells []string `json:"cells"`
}

func ProcessExcel(w http.ResponseWriter, r *http.Request) {
	// Get the uploaded file from the request
	// Setting Up the CORS header so that anybody from any port or host can hit the api
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Process Excel Request received")

	file, header, err := r.FormFile("excelFile")
	if err != nil {
		http.Error(w, "Error getting file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("File Name: %s, Size: %d\n", header.Filename, header.Size)
	xlFile, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, "Error reading Excel file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var rows []ExcelRow

	// Iterate through the rows and build the response data
	rowsLimit := 5 // Change this to the desired number of rows to process

	sheetName := xlFile.GetSheetName(0)

	rowsData, err := xlFile.GetRows(sheetName)
	if err != nil {
		http.Error(w, "Error getting rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for rowIndex, row := range rowsData {
		if rowIndex >= rowsLimit {
			break
		}

		rowCells := append([]string{}, row...)

		rows = append(rows, ExcelRow{Cells: rowCells})
	}

	// Marshal the response data into JSON
	jsonData, err := json.Marshal(rows)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

// Inital work done with this library
//	"github.com/tealeg/xlsx"

// func ProcessExcel(w http.ResponseWriter, r *http.Request) {
// 	// Get the uploaded file from the request
// 	// Setting Up the CORS header so that anybody from any port or host can hit the api
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	fmt.Println("Process Excel Request received")
// 	file, _, err := r.FormFile("excelFile")
// 	if err != nil {
// 		http.Error(w, "Error getting file", http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	file, header, err := r.FormFile("excelFile")
// 	if err != nil {
// 		http.Error(w, "Error getting file: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	fmt.Printf("File Name: %s, Size: %d\n", header.Filename, header.Size)

// 	// Read and print the first 5 rows from the Excel file
// 	xlFile, err := xlsx.OpenReaderAt(file, header.Size)
// 	if err != nil {
// 		http.Error(w, "Error reading Excel file: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	var rows []ExcelRow

// 	// Iterate through the rows and build the response data

// 	for _, row := range xlFile.Sheets[0].Rows[:5] {
// 		var rowCells []string
// 		for _, cell := range row.Cells {
// 			rowCells = append(rowCells, cell.String())
// 		}
// 		rows = append(rows, ExcelRow{Cells: rowCells})
// 	}

// 	// Marshal the response data into JSON
// 	jsonData, err := json.Marshal(rows)
// 	if err != nil {
// 		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Write the JSON response
// 	w.Write(jsonData)

// }
