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

	apiRouter := root.Group("bing")
	{
		// bing的壁纸获取
		apiRouter.GET("wallpaper", bing.GetWallpaper)
	}
}
