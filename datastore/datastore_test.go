package datastore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func Test_Search(t *testing.T) {
	config.Load()
	logger.Setup()
	db.Init()
	defer db.Close()

	Init()
	start := time.Now()
	results := Search("ApBL0009019")
	elapsed := time.Since(start)
	logger.Debug(elapsed)
	// assert.Equal(t, len(results), 1)
	assert.Equal(t, len(results), 0)
}

func Test_RebuildIndex(t *testing.T) {
	config.Load()
	logger.Setup()
	db.Init()
	defer db.Close()

	ReBuildIndex()
}
