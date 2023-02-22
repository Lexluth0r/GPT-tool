package response

import (
	"chat/app/common/errcode"
	"github.com/gin-gonic/gin"

	"net/http"
)

// 成功响应
func Success(c *gin.Context, data ...interface{}) {
	response := gin.H{
		"code": errcode.SUCCESS,
		"msg":  errcode.GetMsg(errcode.SUCCESS),
	}

	if len(data) == 0 {
		response["data"] = [0]string{}
		c.JSONP(http.StatusOK, response)
		return
	}
	if len(data) == 1 {
		response["data"] = data[0]
	}

	c.JSONP(http.StatusOK, response)
}

// 自定义错误
func Error(c *gin.Context, globalCode int, data ...interface{}) {

	response := gin.H{
		"code": globalCode,
		"msg":  errcode.GetMsg(globalCode),
	}

	if len(data) == 0 {
		response["data"] = [0]string{}
		c.JSONP(http.StatusOK, response)
		return
	}
	if len(data) == 1 {
		response["data"] = data
	}

	c.JSONP(http.StatusOK, response)

}

//服务异常
func ErrHandle(c *gin.Context) {

	c.JSONP(http.StatusInternalServerError, map[string]interface{}{
		"code": errcode.ERROR,
		"msg":  errcode.GetMsg(errcode.ERROR),
		"data": [0]string{},
	})
}
