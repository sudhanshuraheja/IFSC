package model

import (
	"fmt"
	"strconv"
)

// Branch : struct for the data in branch table
type Branch struct {
	DBId      int64  `db:"id" json:"id"`
	Bank      string `db:"bank" json:"bank"`
	Ifsc      string `db:"ifsc" json:"ifsc"`
	Micr      string `db:"micr" json:"micr"`
	Branch    string `db:"branch" json:"branch"`
	Address   string `db:"address" json:"address"`
	City      string `db:"city" json:"city"`
	District  string `db:"district" json:"district"`
	State     string `db:"state" json:"state"`
	Contact   string `db:"contact" json:"contact"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

// ToString : convert Branch object to String
func (b Branch) ToString() string {
	return fmt.Sprintf("[%s] %s %s %s %s %s %s %s %s %s %s %s", strconv.FormatInt(b.DBId, 10), b.Bank, b.Ifsc, b.Micr, b.Branch, b.Address, b.City, b.District, b.State, b.Contact, b.CreatedAt, b.UpdatedAt)
}
