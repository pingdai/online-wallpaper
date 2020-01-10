package wx

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	// var res = make(map[string]interface{})
	var err error
	var req = &VerifyReq{}
	defer func() {
		c.Writer.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(gin.H{"echostr": req.Echostr})
		c.Writer.Write(b)
		logrus.Infof("微信回调返回：%s", string(b))
	}()

	if err = c.Bind(req); err != nil {
		logrus.Warnf("微信回调 参数错误，req:%+v err:%v", req, err)
		return
	}
	logrus.Infof("微信回调 参数 req:%+v", *req)

	return
}
