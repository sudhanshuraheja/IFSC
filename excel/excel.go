package excel

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// Load : Try reading from the excel
func Load() {
	fmt.Println("Running this now")

	fileName := "../data/sample.xlsx"

	workBook, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Println("There was an error in opening the file:", err)
		return
	}

	for _, sheet := range workBook.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
