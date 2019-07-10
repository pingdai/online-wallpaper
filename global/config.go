package global

import (
	"github.com/pingdai/tools/configwatcherx"
	"github.com/pingdai/tools/ginx"
	"github.com/pingdai/tools/log"
	"github.com/pingdai/tools/servicex"
)

const (
	PROJECT_NAME = "online-wallpaper"
)

var Config Cfg

func init() {
	servicex.SetServiceName(PROJECT_NAME)
	servicex.ConfP(&Config)
	cw.ConfigWatcher.Init()
}

type Cfg struct {
	Log  *log.Log   `json:"log"`
	Ginx *ginx.Ginx `json:"ginx"`
}
