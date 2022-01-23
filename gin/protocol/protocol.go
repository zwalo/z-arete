package protocol

import (
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

type ResultCode int

const (
	SUCCESS ResultCode = 0
	FAILED             = 1
)

func (r ResultCode) String() string {
	switch r {
	case SUCCESS:
		return "Success"
	case FAILED:
		return "Failed"
	}
	return ""
}

type RespHeader struct {
	Result       ResultCode  `json:"result"`
	ResultString string      `json:"resultString"`
	Desc         interface{} `json:"desc"`
	Data         interface{} `json:"data"`
}

func NewErrorResp(err error) *RespHeader {
	resultCode := ResultCode(FAILED)
	return &RespHeader{
		resultCode,
		resultCode.String(),
		err.Error(),
		nil,
	}
}
func NewResp(res interface{}) *RespHeader {
	resultCode := ResultCode(SUCCESS)
	return &RespHeader{
		resultCode,
		resultCode.String(),
		nil,
		res,
	}
}
func RespError(c *gin.Context, status int, err error, logEntry log.Logger) {
	logEntry.Error("error", err, "request", c.Request.URL.Path)
	c.JSON(status, NewErrorResp(err))
}

func RespResponse(c *gin.Context, response interface{}, logEntry log.Logger) {
	logEntry.Trace(c.Request.URL.Path)
	c.JSON(http.StatusOK, NewResp(response))
}
