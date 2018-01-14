package excel_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/excel"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func Test_one(t *testing.T) {
	config.Load()
	logger.Setup()

	dir, _ := os.Getwd()
	logger.Infoln(dir)

	allBranches := excel.Load("../data/sample_tiny.xlsx")
	assert.Equal(t, allBranches.Count, 12)
	assert.NotEmpty(t, allBranches.List[11].Bank)
}
