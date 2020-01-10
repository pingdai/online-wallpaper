package wx

import (
	"github.com/gin-gonic/gin"
	"github.com/pingdai/online-wallpaper/modules"
	"github.com/sirupsen/logrus"
)

// swagger:parameters VerifyReq
type VerifyReq struct {
	// 签名
	//  required:false
	Signature string `json:"signature" form:"signature"`
	// 回显字符串
	// required:true
	Echostr string `json:"echostr" form:"echostr" binding:"required"`
	// 时间戳
	// required:false
	Timestamp string `json:"timestamp" form:"timestamp"`
	//
	// required:false
	Nonce string `json:"nonce" form:"nonce"`
}

// swagger:response VerifyRes
type VerifyRes struct {
	// in:body
	Rsp struct {
		Data struct {
			Echostr string `json:"echostr"`
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
	var req = &VerifyReq{}
	defer func() {
		if err != nil {
			modules.SendResponse(c, -1, err.Error(), res)
		} else {
			modules.SendResponse(c, 0, "succ", res)
		}
	}()

	if err = c.Bind(req); err != nil {
		logrus.Warnf("微信回调 参数错误，req:%+v err:%v", req, err)
		return
	}
	logrus.Infof("微信回调 参数 req:%+v", *req)

	res["echostr"] = req.Echostr

	return
}
