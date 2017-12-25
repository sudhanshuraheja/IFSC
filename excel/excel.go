package excel

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// Branches : Array and Count of all branches that we could find in the excel
type Branches struct {
	count int
	list  []Branch
}

// Branch : list of details of a bank branch
type Branch struct {
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

func (b *Branch) populate(row *xlsx.Row) {
	for index, cell := range row.Cells {
		switch index {
		case 0:
			b.bank = cell.String()
		case 1:
			b.ifsc = cell.String()
		case 2:
			b.micr = cell.String()
		case 3:
			b.branch = cell.String()
		case 4:
			b.address = cell.String()
		case 5:
			b.contact = cell.String()
		case 6:
			b.city = cell.String()
		case 7:
			b.district = cell.String()
		case 8:
			b.state = cell.String()
		default:
			fmt.Println("Mismatcing colums found")
		}
	}
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
			sheetRow := Branch{}
			sheetRow.populate(row)
			allBranches.list = append(allBranches.list, sheetRow)
			allBranches.count++
		}
	}

	return allBranches
}
