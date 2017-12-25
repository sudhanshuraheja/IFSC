package excel

import (
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/tealeg/xlsx"
)

// Branches : Array and Count of all branches that we could find in the excel
type Branches struct {
	Count int
	List  []Branch
}

// Branch : list of details of a bank branch
type Branch struct {
	Bank     string
	Ifsc     string
	Micr     string
	Branch   string
	Address  string
	Contact  string
	City     string
	District string
	State    string
}

func (b *Branch) populate(row *xlsx.Row) {
	for index, cell := range row.Cells {
		switch index {
		case 0:
			b.Bank = cell.String()
		case 1:
			b.Ifsc = cell.String()
		case 2:
			b.Micr = cell.String()
		case 3:
			b.Branch = cell.String()
		case 4:
			b.Address = cell.String()
		case 5:
			b.Contact = cell.String()
		case 6:
			b.City = cell.String()
		case 7:
			b.District = cell.String()
		case 8:
			b.State = cell.String()
		default:
			logger.Error("Mismatcing colums found")
		}
	}
}

// Load : Try reading from the excel
func Load(file string) Branches {
	allBranches := Branches{}

	workBook, err := xlsx.OpenFile(file)
	if err != nil {
		logger.Debug(err.Error())
		return allBranches
	}

	for _, sheet := range workBook.Sheets {
		for _, row := range sheet.Rows {
			sheetRow := Branch{}
			sheetRow.populate(row)
			allBranches.List = append(allBranches.List, sheetRow)
			allBranches.Count++
		}
	}

	return allBranches
}
