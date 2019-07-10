package cw

import (
	"fmt"
	"github.com/pingdai/tools/constants"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type WatcherHandler func([]byte, EventType, error)

var ConfigWatcher = new(ConfigWatcherX)

// zk config watcher
type ConfigWatcherX struct {
	WatcherHandler WatcherHandler

	c    *zk.Conn
	init bool
}

func (cwx *ConfigWatcherX) Init() {
	if !cwx.init {
		cwx.New()
		cwx.init = true
	}
}

func (cwx *ConfigWatcherX) MarshalDefaults() {
	// 如果是从本地起的配置文件，则不需要监听了
	confSource := os.Getenv(constants.I_EnvVarKeyConfigSource)
	if confSource == "1" { // local
		cwx.init = true
		return
	}

	configAddress := os.Getenv(constants.O_EnvVarKeyConfigAddress)
	if configAddress == "" {
		panic(fmt.Sprintf("%s cannot be empty,usage: 127.0.0.1:8080,127.0.0.1:8081", constants.O_EnvVarKeyConfigAddress))
	}

	configPath := os.Getenv(constants.I_EnvVarKeyConfigSource)
	if configPath == "" {
		panic(fmt.Sprintf("%s cannot be empty,usage: /entry/config/service/%s", constants.I_EnvVarKeyConfigSource, constants.EnvVarKeyProjectName))
	}

	zk.DefaultLogger = logrus.StandardLogger()
}

func (cwx *ConfigWatcherX) New() {
	cwx.MarshalDefaults()
	if cwx.init {
		return
	}

	var err error
	// 注册
	node := os.Getenv(constants.O_EnvVarKeyConfigAddress)
	servers := strings.Split(node, ",")
	cwx.c, _, err = zk.Connect(servers, 3*time.Second)
	if err != nil {
		panic(fmt.Sprintf("Connect zookeeper fail,err:%v", err))
	}
	// defer c.Close()

	configPath := os.Getenv(constants.I_EnvVarKeyZKConfigPath)

	existed, err := Existed(cwx.c, configPath)
	if err != nil {
		panic(fmt.Sprintf("从zk判断路径是否存在，失败。zk_node[%s] path[%s] err:%v",
			node, configPath, err))
	}
	if !existed {
		panic(fmt.Sprintf("zk path[%s] not exist", configPath))
	}

	// 进行文件内容监控
	go cwx.watchContentChange()
}

// 只监听文件内容变动
// 有报错立即停止监听
func (cwx *ConfigWatcherX) watchContentChange() {
	if cwx.WatcherHandler == nil {
		logrus.Warnf("未设置配置自动监听回调方法，自动监听将不生效")
		return
	}

	configPath := os.Getenv(constants.I_EnvVarKeyZKConfigPath)
	logrus.Infof("配置监听开始")
	for {
		_, _, event, err := cwx.c.GetW(configPath)
		if err != nil {
			cwx.WatcherHandler([]byte{}, EventNodeUnknown, err)
			logrus.Errorf("配置监听失败，准备退出监听模式，err:%v", err)
			break
		}

		evt := <-event

		if evt.Type == zk.EventNodeDeleted { // 配置节点被删除
			cwx.WatcherHandler([]byte{}, EventNodeDeleted, nil)
			logrus.Errorf(fmt.Sprintf("配置监听到路径[%s]被删除，准备退出监听模式", configPath))
			break
		} else if evt.Type == zk.EventNodeDataChanged { // 内容变更
			logrus.Infof("配置监听到内容变更")
			content, _, err := cwx.c.Get(configPath)
			if err != nil {
				cwx.WatcherHandler([]byte{}, EventNodeUnknown, err)
				logrus.Errorf("配置监听获取失败，准备退出监听模式，err:%v", err)
				break
			}

			cwx.WatcherHandler(content, EventNodeDataChanged, nil)
		} else {
			logrus.Infof("配置监听的事件：%s，不回调", evt.Type.String())
		}
	}
}
