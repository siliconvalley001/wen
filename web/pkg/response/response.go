package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const CtxNickNameKey = "nickname"

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code int, msg string) *Response {

	res := &Response{Code: code, Msg: msg}
	c.JSON(http.StatusInternalServerError,res)
	c.Set("response", res)
	return res
}

func ResponseSucess(c *gin.Context, code int, data interface{}) *Response {
	res := &Response{Code: code, Msg: data}
	c.JSON(200, res)
	res2, _ := json.Marshal(res)
	c.Set("response", res2)
	return res
}
