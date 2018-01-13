package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
)

func Test_start(t *testing.T) {
	config.Load()

	app := Start()
	assert.Equal(t, app.Usage, "this service lists all bank branches and ifsc codes in india")
}
