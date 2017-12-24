package excel

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// Branches : Array and Count of all branches that we could find in the excel
type Branches struct {
	count int
	list  []branch
}

type branch struct {
	bank     string
	ifsc     string
	micr     string
	branch   string
	address  string
	contact  string
	city     string
	district string
	state    string
}

// Load : Try reading from the excel
func Load(file string) Branches {
	allBranches := Branches{}

	workBook, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println("There was an error in opening the file:", err)
		return allBranches
	}

	for _, sheet := range workBook.Sheets {
		for _, row := range sheet.Rows {
			sheetRow := populateCell(row)
			allBranches.list = append(allBranches.list, sheetRow)
			allBranches.count++
		}
	}

	return allBranches
}

func populateCell(row *xlsx.Row) branch {
	sheetRow := branch{}
	for index, cell := range row.Cells {
		switch index {
		case 0:
			sheetRow.bank = cell.String()
		case 1:
			sheetRow.ifsc = cell.String()
		case 2:
			sheetRow.micr = cell.String()
		case 3:
			sheetRow.branch = cell.String()
		case 4:
			sheetRow.address = cell.String()
		case 5:
			sheetRow.contact = cell.String()
		case 6:
			sheetRow.city = cell.String()
		case 7:
			sheetRow.district = cell.String()
		case 8:
			sheetRow.state = cell.String()
		default:
			fmt.Println("Mismatcing colums found")
		}
	}
	return sheetRow
}
