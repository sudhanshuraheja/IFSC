package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_one(t *testing.T) {
	allBranches := Load("../data/sample.xlsx")
	assert.Equal(t, allBranches.count, 139491)
}
