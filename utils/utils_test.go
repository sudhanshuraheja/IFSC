package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/utils"
)

func Test_mergeMaps(t *testing.T) {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	map2 := map[string]int{"two": 2, "three": 3, "four": 4}
	assert.Equal(t, utils.MergeMaps(map1, map2), map[string]int{"one": 1, "two": 4, "three": 6, "four": 4})
}
