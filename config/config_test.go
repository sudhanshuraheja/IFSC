package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
)

func Test_ConfigValues(t *testing.T) {
	config.Load()
	assert.Equal(t, "ifsc", config.Name())

	assert.Equal(t, "dbname=ifsc user=ifsc password='ifsc_postgres' host=db port=5432 sslmode=disable", config.Database().ConnectionString())
	assert.Equal(t, "postgres://ifsc:ifsc_postgres@db:5432/ifsc?sslmode=disable", config.Database().ConnectionURL())
}
