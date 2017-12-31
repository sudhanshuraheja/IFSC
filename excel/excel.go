package excel

import (
	"strings"

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
	Bank     string `json:"bank"`
	Ifsc     string `json:"ifsc"`
	Micr     string `json:"micr"`
	Branch   string `json:"branch"`
	Address  string `json:"address"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	Contact  string `json:"contact"`
}

func (b *Branch) populate(row *xlsx.Row) {
	for index, cell := range row.Cells {
		textEntry := formatString(cell.String())
		switch index {
		case 0:
			b.Bank = textEntry
		case 1:
			b.Ifsc = textEntry
		case 2:
			b.Micr = textEntry
		case 3:
			b.Branch = textEntry
		case 4:
			b.Address = textEntry
		case 5:
			b.Contact = textEntry
		case 6:
			b.City = textEntry
		case 7:
			b.District = textEntry
		case 8:
			b.State = textEntry
		default:
			logger.Error("Mismatcing colums found")
		}
	}
}

// Load : Try reading from the excel
func Load(file string) Branches {
	allBranches := Branches{}

	logger.Infoln("Going to start reading file", file)
	workBook, err := xlsx.OpenFile(file)
	if err != nil {
		logger.Debug(err.Error())
		return allBranches
	}

	for sheetNumber, sheet := range workBook.Sheets {
		logger.Infoln("Reading sheet", sheetNumber)
		for _, row := range sheet.Rows {
			sheetRow := Branch{}
			sheetRow.populate(row)
			allBranches.List = append(allBranches.List, sheetRow)
			allBranches.Count++
		}
	}

	return allBranches
}

func formatString(text string) string {
	textEntry := strings.ToLower(text)

	switch textEntry {
	case "na":
		textEntry = "NA"
	default:
		textEntry = strings.Title(textEntry)
	}

	return textEntry
}
