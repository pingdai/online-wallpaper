package main

import (
	"github.com/pingdai/online-wallpaper/global"
	"github.com/pingdai/online-wallpaper/routes"
)

func main() {
	routes.RootRouter(global.Config.Ginx.Engine)
	global.Config.Ginx.Run()
}
