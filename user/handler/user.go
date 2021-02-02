package handler

import (
	"context"
	"github.com/siliconvalley001/wen/user/model"
	ex "github.com/siliconvalley001/wen/user/proto"
	"github.com/siliconvalley001/wen/user/service"
)


func (h *Handler)Register(context context.Context, req *ex.ResuqestRegister,resp *ex.ResponseRegister) error {
	userRegister:=&model.User{
		NickName: req.NickName,
		Password: req.Password,
		Avatar: req.Avatar,
		Name: req.Name,
	}
	_,err:=service.AddUser(userRegister)
	if err!=nil{
		return err
	}
	resp.Msg="用户添加成功"
	return nil
}

func (h *Handler)Login(context context.Context, req *ex.ResquestLogin, resp *ex.ResponseLogin) error {
	err:=service.CheckPassWord(req.NickName,req.Password)
	if err!=nil{
		return err
	}
	resp.Msg="用户登陆成功"
	return nil

}

func (h *Handler)GetUserInfo(context context.Context, req *ex.ResquestUserInfo, resp *ex.ResponseUserInfo) error {
	userModel,err:=service.FindUserByNickName(req.NickName)
	if err!=nil{
		return err
	}
	UserToUserInfo(userModel)
	return nil

}

func UserToUserInfo(user *model.User)*ex.ResponseUserInfo{
	return &ex.ResponseUserInfo{}
}




