package user

import "time"

type UserLoginResquest struct {
	NickName string `json:"nick_name" form:"nick_name"`
	Password string `json:"password" form:"password"`
}
type UserLoginResponse struct {
	Code      int32     `json:"code"`
	Msg       string    `json:"msg"`
	NickName  string    `json:"nick_name"`
	LoginTime time.Time `json:"login_time"`
	Token     string    `json:"token"`
}
type UserRegisterRequest struct {
	Avatar     string `json:"avatar"`
	NickName   string `json:"nick_name"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Sex        int32  `json:"sex"`
	Tel        string `json:"tel"`
}
type UserRegisterResponse struct {
}
