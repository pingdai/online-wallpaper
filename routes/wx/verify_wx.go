package wx

import (
	"github.com/gin-gonic/gin"
	"github.com/pingdai/online-wallpaper/modules"
)

// swagger:parameters VerifyReq
type VerifyReq struct {
	// 请求图片截止天数，0今天，-1截止明天，1截止昨天（以此类推，目前最多获取到七天前的图片）
	//  required:false
	Idx int `json:"idx" form:"idx"`
	// 1~8请求返回数量，目前最多一次获取8张
	// required:true
	N int `json:"n" form:"n" binding:"required"`
}

// swagger:response VerifyRes
type VerifyRes struct {
	// in:body
	Rsp struct {
		Data struct {
		} `json:"data"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
}

// swagger:route GET /online-wallpaper/wx/verify WX VerifyReq
//
// 验证微信
//
// 验证微信
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
//       200: VerifyRes
//		 default: CommRes
func Verify(c *gin.Context) {
	var res = make(map[string]interface{})
	var err error
	// var req = &VerifyReq{}
	defer func() {
		if err != nil {
			modules.SendResponse(c, -1, err.Error(), res)
		} else {
			modules.SendResponse(c, 0, "succ", res)
		}
	}()

	// if err = c.Bind(req); err != nil {
	// 	logrus.Warnf("获取必应壁纸 参数错误，req:%+v err:%v", req, err)
	// 	return
	// }
	// logrus.Infof("获取必应壁纸 参数 req:%+v", req)

	return
}
