package bing

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pingdai/online-wallpaper/modules"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// swagger:parameters GetPermissionSelfListReq
type GetWallpaperReq struct {
	// 请求图片截止天数，0今天，-1截止明天，1截止昨天（以此类推，目前最多获取到七天前的图片）
	//  required:false
	Idx int `json:"idx" form:"idx"`
	// 1~8请求返回数量，目前最多一次获取8张
	// required:true
	N int `json:"n" form:"n" binding:"required"`
}

// swagger:response GetPermissionSelfListRes
type GetWallpaperRes struct {
	// in:body
	Rsp struct {
		Data struct {
			Results []BingWallpaperItem `json:"results"`
		} `json:"data"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
}

type BingWallpaperItem struct {
	// 开始时间
	Startdate string `json:"startdate"`
	// 结束时间
	Enddate string `json:"enddate"`
	// 访问链接
	Url string `json:"url"`
	// 访问原始链接
	Urlbase string `json:"urlbase"`
	//
	Copyright string `json:"copyright"`
	//
	Copyrightlink string `json:"copyrightlink"`
}

type BindWallpaperImgs struct {
	Images []BingWallpaperItem `json:"images"`
}

// swagger:route GET /online-wallpaper/bind/wallpaper Bing GetWallpaperReq
//
// 获取必应壁纸
//
// 获取必应壁纸
//
//     Consumes:
//     - application/x-www-form-urlencoded
//	   - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: GetPermissionSelfListRes
//		 default: CommRes
func GetWallpaper(c *gin.Context) {
	var res = make(map[string]interface{})
	var err error
	var req = &GetWallpaperReq{}
	defer func() {
		if err != nil {
			modules.SendResponse(c, -1, err.Error(), res)
		} else {
			modules.SendResponse(c, 0, "succ", res)
		}
	}()

	if err = c.Bind(req); err != nil {
		logrus.Warnf("获取必应壁纸 参数错误，req:%+v err:%v", req, err)
		return
	}
	logrus.Infof("获取必应壁纸 参数 req:%+v", req)

	url := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=%d&mkt=de-DE", req.Idx, req.N)
	logrus.Debugf("url:%s", url)
	rsp, err := http.Get(url)
	if err != nil {
		logrus.Warnf("request [%s] err:%v", url, err)
		return
	}
	defer rsp.Body.Close()

	content, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		logrus.Warnf("read [%s] rsp content err:%v", url, err)
		return
	}

	var bindWallpaperImgs BindWallpaperImgs
	if err = json.Unmarshal(content, &bindWallpaperImgs); err != nil {
		logrus.Warnf("Unmarshal content[%s] err:%v", string(content), err)
		return
	}

	list := make([]BingWallpaperItem, 0)
	for _, v := range bindWallpaperImgs.Images {
		v.Url = fmt.Sprintf("https://www.bing.com%s", v.Url)
		v.Urlbase = fmt.Sprintf("https://www.bing.com%s", v.Urlbase)
		list = append(list, v)
	}

	res["results"] = list

	return
}
