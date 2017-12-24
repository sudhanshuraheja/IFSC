package excel

import (
	"fmt"

	"github.com/extrame/xls"
)

func Load() {
	if xlFile, err := xls.Open("Table.xls", "utf-8"); err == nil {
		fmt.Println(xlFile.Author)
	}
}
