package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
)

func Test_ConfigValues(t *testing.T) {
	config.Load()
	assert.Equal(t, "ifsc", config.Name())
	assert.Equal(t, "0.0.1", config.Version())
	assert.Equal(t, "debug", config.LogLevel())
	assert.Equal(t, "3000", config.Port())
	assert.Equal(t, false, config.EnableStaticFileServer())
	assert.Equal(t, true, config.EnableGzipCompression())
	assert.Equal(t, false, config.EnableDelayMiddleware())
	assert.Equal(t, "file://./db/sample.xlsx", config.LatestDataExcel())
	assert.Equal(t, false, config.MapsEnabled())
	assert.Equal(t, "google-maps-api-key", config.MapsKey())

	assert.Equal(t, "dbname=ifsc user=ifsc password='ifsc_postgres' host=db port=5432 sslmode=disable", config.Database().ConnectionString())
	assert.Equal(t, "postgres://ifsc:ifsc_postgres@db:5432/ifsc?sslmode=disable", config.Database().ConnectionURL())
	assert.Equal(t, 10, config.Database().MaxPoolSize())
}
