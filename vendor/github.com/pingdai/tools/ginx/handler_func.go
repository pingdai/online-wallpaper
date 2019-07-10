package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pingdai/tools/constants"
	"github.com/pingdai/tools/duration"
	logContext "github.com/pingdai/tools/log/context"
	"github.com/pingdai/tools/log/hooks/merchant_code"
	"github.com/pingdai/tools/str"
	"github.com/sirupsen/logrus"
)

/*
	存放的是gin里面的一些中间件
*/

// Gin请求日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		d := duration.NewDuration()

		method := c.Request.Method
		statusCode := c.Writer.Status()
		UA := c.Request.UserAgent()

		reqID := c.Request.Header.Get(constants.HEADER_LOD_ID)
		if reqID == "" {
			reqID = uuid.New().String()
		}
		logContext.SetLogID(reqID)
		//p := c.Request.Header.Get(constants.HEADER_PHONE)
		//phone.SetPhone(p)
		//openID := c.Request.Header.Get(constants.HEADER_OPEN_ID)
		//open_id.SetOpenID(openID)
		//userID := c.Request.Header.Get(constants.HEADER_USER_ID)
		//user_id.SetUserID(userID)
		//ui := c.Request.Header.Get(constants.HEADER_UIN)
		//uin.SetUin(ui)
		rs := c.Request.Header.Get(constants.HEADER_REMOTE_SERVICE)
		mc := c.Request.Header.Get(constants.HEADER_MERCHANT_CODE)
		merchant_code.SetMerchantCode(mc)

		requestData := GetRequestData(c)
		urlPath := fmt.Sprintf("%s%s", c.Request.Host, str.Cuts(requestData, 2048))

		// Process request
		c.Next()

		fields := logrus.Fields{
			"cost_time":      d.Get(),
			"remote_service": rs,
			"method":         method,
			"status":         statusCode,
			"user_agent":     UA,
			"url_path":       urlPath,
			"merchant_code":  mc,
		}

		logger := logrus.WithFields(fields)

		logger.Infof("")
	}
}

func GetRequestData(c *gin.Context) string {
	var requestData string
	method := c.Request.Method
	if method == "GET" || method == "DELETE" {
		requestData = c.Request.RequestURI
	} else {
		c.Request.ParseForm()
		requestData = fmt.Sprintf("%s [%s]", c.Request.RequestURI, c.Request.Form.Encode())
	}
	return requestData
}
