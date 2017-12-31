package datastore

import (
	"github.com/sudhanshuraheja/ifsc/excel"
)

// Branch : struct for the data in branch table
type Branch struct {
	DBId      int64  `db:"id"`
	Bank      string `db:"bank"`
	Ifsc      string `db:"ifsc"`
	Micr      string `db:"micr"`
	Branch    string `db:"branch"`
	Address   string `db:"address"`
	City      string `db:"city"`
	District  string `db:"district"`
	State     string `db:"state"`
	Contact   string `db:"contact"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// Search : search the datastore
func Search(query string) []excel.Branch {
	results := []excel.Branch{}
	return results
}
