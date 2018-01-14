package excel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func Test_download(t *testing.T) {
	config.Init()
	logger.Init()
	err := download("http://localhost:"+config.Port()+"/ping", "ping.html")
	assert.Equal(t, nil, err)

	os.Remove("ping.html")
}
