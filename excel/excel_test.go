package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_one(t *testing.T) {
	Load()
	assert.Equal(t, 1, 1)
}
