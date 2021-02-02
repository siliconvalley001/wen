package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"time"
	"github.com/siliconvalley001/wen/web/model/user"
	"github.com/siliconvalley001/wen/web/pkg/jwt"
	"github.com/siliconvalley001/wen/web/pkg/response"
	"fmt"

	//"github.com/micro/go-micro/v2/server/grpc"

	//"github.com/micro/go-micro/v2/server/grpc"
	//"github.com/micro/go-micro/v2"
	user2 "github.com/siliconvalley001/wen/user/proto"
	//"github.com/micro/go-plugins/transport/grpc"
	//"github.com/micro/micro/service/server/grpc"
)
//@Summary 用户登陆
//@Description 用户登陆
//@Tags 用户登陆接口
//@ID /user/login
//@Accept json
//@Produce json
//@Param body body user.UserLoginResquest true "body"
//@Success 200  {object}   response.Response{data=user.UserLoginResponse}  "success"
//@Router /user/login [post]
func UserLogin(c *gin.Context)  {
	//使用grpc通信
	param:=new(user.UserLoginResquest)
	if err:=c.ShouldBind(param);err!=nil{
		response.ResponseError(c,40001,"登陆参数绑定失败"+err.Error())
		return
	}
	srv:=grpc.NewService()
	US:=user2.NewUserService("micro.user",srv.Client())
	resquest:=&user2.ResquestLogin{
		NickName: param.NickName,
		Password: param.Password,
	}
	response1,err:=US.Login(c,resquest)
	if err!=nil{
		response.ResponseError(c,40002,"登陆失败"+err.Error())
		return
	}

	token,err:= jwt.GenToken1(response1.NickName)
	if err!=nil{
		response.ResponseError(c,40003,"token生成失败"+err.Error())
		return
	}

	sessInfo:=&user.UserLoginResponse{
		Code: 200,
		Msg: "登陆成功",
		LoginTime: time.Now(),
		NickName: param.NickName,
		Token:token,
	}
	c.Set(response.CtxNickNameKey,param.NickName)
	response.ResponseSucess(c,200,sessInfo)
}

//@Summary 用户注册
//@Description 用户注册
//@Tags 用户注册接口
//@ID /user/register
//@Accept json
//@Produce json
//@Param body body user.UserRegisterRequest true "body"
//@Success 200  {object}   response.Response{data=user.UserRegisterResponse}  "success"
//@Router /user/register [post]
func UserRegister(c *gin.Context)  {

	param:=new(user.UserRegisterRequest)
	if err:=c.ShouldBind(param);err!=nil{
		response.ResponseError(c,40001,"注册参数绑定失败"+err.Error())
		return
	}
	srv:=grpc.NewService()
	fmt.Println(param)
	registerService:=user2.NewUserService("micro.user",srv.Client())
	inresquest:=user2.ResuqestRegister{
		Avatar: param.Avatar,
		NickName: param.NickName,
		Password: param.Password,
		RePassword: param.RePassword,
		Sex: param.Sex,
		Tel: param.Tel,
	}
	fmt.Println(inresquest)
	responseRegister,err:=registerService.Register(context.TODO(),&inresquest)
	if err!=nil{
		response.ResponseError(c,40002,"注册失败"+err.Error())
		fmt.Println(err.Error())

		return
	}
	response.ResponseSucess(c,200,responseRegister)
}
