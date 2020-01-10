package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pingdai/online-wallpaper/routes/bing"

	"github.com/pingdai/online-wallpaper/global"
	"github.com/pingdai/tools/courier/checkhealth"
	"github.com/pingdai/tools/courier/swagger"
)

func RootRouter(engine *gin.Engine) {
	root := engine.Group(global.PROJECT_NAME)
	{
		swagger.Init(root)
		checkhealth.Init(root)
	}

	bingRouter := root.Group("bing")
	{
		// bing的壁纸获取
		bingRouter.GET("wallpaper", bing.GetWallpaper)
	}

	wxRouter := root.Group("wx")
	{
		// 接入微信
		wxRouter.GET("verify", wx.Verify)
	}
}
