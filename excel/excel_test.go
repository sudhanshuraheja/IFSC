package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func Test_one(t *testing.T) {
	config.Load()
	logger.Setup()

	allBranches := Load("../data/sample.xlsx")
	assert.Equal(t, allBranches.count, 139491)
	assert.NotEmpty(t, allBranches.list[139490].bank)
}
