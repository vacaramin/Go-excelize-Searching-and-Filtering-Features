package models

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelRow struct {
	Cells []string `json:"cells"`
}
type Filter struct {
	Column int
	Row    int
}

func (f *Filter) ColumnInAlphabets() string {
	x, _ := excelize.ColumnNumberToName(f.Column)
	return x
}

func (f *Filter) CellAddress() string {
	x, _ := excelize.ColumnNumberToName(f.Column)
	return x + strconv.FormatInt(int64(f.Row), 10)
}
