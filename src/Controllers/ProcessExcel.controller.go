package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	handlers "xlsx/src/Handlers"
	models "xlsx/src/Models"

	"github.com/xuri/excelize/v2"
)

var (
	F   models.Filter
	FM  models.Filter
	FV  models.Filter
	G1Y models.Filter
	G3N models.Filter
	C1N models.Filter
	C2O models.Filter
	C3S models.Filter
)

// ProcessExcel This is the main function that takes the excel and implements all the required logic on the excel and returns an searched and filtered excel
func ProcessExcel(w http.ResponseWriter, r *http.Request) {
	// Setting Up the CORS header so that anybody from any port or host can hit the api
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Process Excel Request received")

	file, header, err := r.FormFile("excelFile")
	if err != nil {
		handlers.ErrorHandler(w, http.StatusBadRequest, "Error getting file "+err.Error())
		return
	}
	defer file.Close()

	fmt.Printf("File Name: %s, Size: %d\n", header.Filename, header.Size)
	xlFile, err := excelize.OpenReader(file)
	if err != nil {
		handlers.ErrorHandler(w, http.StatusInternalServerError, "Error reading Excel file:"+err.Error())
		return
	}
	if xlFile.SheetCount > 1 {
		handlers.ErrorHandler(w, http.StatusBadRequest, "The File has more than 1 sheet, Can't process file")
		return
	}
	var rows []models.ExcelRow

	// Iterate through the rows and build the response data
	rowsLimit := 5 // Change this to the desired number of rows to process

	sheetName := xlFile.GetSheetName(0)
	// Checking First column and it's filters
	for i, j := 1, ""; i <= 25; i++ {
		j = strconv.FormatInt(int64(i), 10)
		cell, err := xlFile.GetCellValue(sheetName, "A"+j)
		if err != nil {
			log.Println("hello")
		}
		if cell != "" {
			log.Print(cell)
		}
		if cell == "#F#" {
			F.Column = 1
			F.Row = i
		}

		if cell != "" && cell[0] != '#' {
			handlers.ErrorHandler(w, http.StatusBadRequest, "The first column has a value that is not a list command, Please check the file")
			return
		}

	}
	fmt.Println(" models.Filter Row = ", F)
	// Checking First Row and it's filters
	for i := 1; i <= 35; i++ {
		j, _ := excelize.ColumnNumberToName(i)
		cell, err := xlFile.GetCellValue(sheetName, j+"1")
		if err != nil {
			log.Println("hello")
		}
		if cell != "" {
			log.Print(cell)
			if cell == "#FM#" {
				FM.Column = i
				FM.Row = 1
			}
			if cell == "#FV#" {
				FV.Column = i
				FV.Row = 1
			}
			if cell == "#G1Y#" {
				G1Y.Column = i
				G1Y.Row = 1
			}
			if cell == "#G3N#" {
				G3N.Column = i
				G3N.Row = 1
			}
			if cell == "#C1N#" {
				C1N.Column = i
				C1N.Row = 1
			}
			if cell == "#C2O#" {
				C2O.Column = i
				C1N.Row = 1
			}
			if cell == "#C3S#" {
				C3S.Column = i
				C3S.Row = 1
			}
		}

		if cell != "" && cell[0] != '#' {
			handlers.ErrorHandler(w, http.StatusBadRequest, "The first Row has a value that is not a list command, Please check the file")
			return
		}

	}
	fmt.Println(" models.Filter Row = ", FM)
	fmt.Println(" FM= ", FM)
	fmt.Println(" FV = ", FV)
	fmt.Println(" G1Y = ", G1Y)
	fmt.Println(" G3N = ", G3N)
	fmt.Println(" C1N = ", C1N)
	fmt.Println(" C2O = ", C2O)
	fmt.Println(" C3S = ", C3S)

	rowsData, err := xlFile.GetRows(sheetName)
	if err != nil {
		handlers.ErrorHandler(w, http.StatusInternalServerError, "Error getting rows: ")
		return
	}

	for rowIndex, row := range rowsData {
		if rowIndex >= rowsLimit {
			break
		}

		rowCells := append([]string{}, row...)

		rows = append(rows, models.ExcelRow{Cells: rowCells})
	}

	// Marshal the response data into JSON
	jsonData, err := json.Marshal(rows)
	if err != nil {
		handlers.ErrorHandler(w, http.StatusInternalServerError, "Error encoding JSON")
		return
	}
	log.Println("Process-excel complete")

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
