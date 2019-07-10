package modules

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Defined sub modules here
type CommonResHead struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

// swagger:response CommRes
type CommRes struct {
	// in:body
	Rsp struct {
		Data struct{} `json:"data"`
		// 0成功，-1失败
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
}

func SendResponse(c *gin.Context, code int, msg string, data interface{}) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	resp := CommonResHead{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("Marshal json to bytes error :%v", err)
	}
	c.Writer.Write(b)

	logrus.Infof("Out: %s", string(b))

	return err
}
