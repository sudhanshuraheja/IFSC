package excel

import (
	"fmt"

	"github.com/extrame/xls"
)

// Load : Try reading from the excel
func Load() {
	fmt.Println("Running this now")

	fileName := "../data/sample.xls"
	encoding := "utf-8"

	workBook, err := xls.Open(fileName, encoding)
	if err != nil {
		fmt.Println("There was an error in opening the file:", err)
		return
	}

	sheetNo := -1
	for {
		sheetNo++
		sheet := workBook.GetSheet(sheetNo)
		if sheet == nil {
			break
		}

		fmt.Println("Found a new sheet [", sheet.Name, "], Total lines: ", sheet.MaxRow)
		// 2 should be int(sheet.MaxRow)
		for i := 1; i < 10; i++ {
			row := sheet.Row(i)
			if row == nil {
				continue
			}

			for col := row.FirstCol(); col < row.LastCol(); col++ {
				fmt.Printf("%s, ", row.Col(col))
			}
			fmt.Println("\n")
		}
	}
}
