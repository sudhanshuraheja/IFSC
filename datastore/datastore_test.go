package datastore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func Test_Search(t *testing.T) {
	config.Load()
	logger.Setup()

	Init("../data/sample.xlsx")
	start := time.Now()
	results := Search("ApBL0009019")
	elapsed := time.Since(start)
	logger.Debug(elapsed)
	assert.Equal(t, len(results), 1)
}
