package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConfigValues(t *testing.T) {
	Load()
	assert.Equal(t, "ifsc", Name())
	assert.Equal(t, "dbname=ifsc user=ifsc password='ifsc_postgres' host=db port=5432 sslmode=disable", config.database.ConnectionString())
	assert.Equal(t, "postgres://ifsc:ifsc_postgres@db:5432/ifsc?sslmode=disable", config.database.ConnectionURL())
}
