package global

import (
	"github.com/pingdai/tools/configwatcherx"
)

func init() {
	cw.ConfigWatcher.WatcherHandler = ConfigWatcherHandler
}

// 服务配置动态监听回调
// eventType:0 - Unknown
// eventType:2 - EventNodeDeleted
// eventType:3 - EventNodeDataChanged
// 先判断err是否为空
func ConfigWatcherHandler(changeContent []byte, eventType cw.EventType, err error) {
	// todo something
}
