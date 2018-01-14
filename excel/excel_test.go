package excel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/excel"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func Test_Load(t *testing.T) {
	config.Init()
	logger.Init()

	allBranches := excel.Load("../data/sample_tiny.xlsx")
	assert.Equal(t, len(allBranches), 12)
	assert.NotEmpty(t, allBranches[11].Bank)
}
