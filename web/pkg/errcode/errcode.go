package errcode

import "fmt"

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}
var codes map[int]string

func NewError(code int,msg string)*Error  {
	if _,ok:=codes[code];ok{
		panic(fmt.Sprintf("错误码 %d 已存在请更换一个",code))
	}
	codes[code]=msg
	return &Error{
		code: code,
		msg: msg,
	}
}

func (e *Error)String()string{
	return fmt.Sprintf("错误码为 %d  错误信息为 %s",e.code,e.msg)
}

func (e *Error)Code()int{
	return e.code
}
func (e *Error)Msg()string{
	return e.msg
}
