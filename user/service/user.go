package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/siliconvalley001/wen/user/dao"
	"github.com/siliconvalley001/wen/user/model"
)

func  AddUser(user *model.User) (id int64, err error) {
	user.Password=generatePwd(user.Password)
	return dao.CreateUser(user)

}

func  DelUser(userid int64) (err error) {
	return dao.DelUserByID(userid)

}

func  UpdateUser(user *model.User, IsChangePWD bool) (err error) {
	if IsChangePWD{
		user.Password=generatePwd(user.Password)

	}
	return dao.UpdateUser(user)
}

func  FindUserByNickName(nickname string) (user *model.User, err error) {

	return dao.FindUserByNickName(nickname)
}

func  CheckPassWord(nickname string,password string)(err error) {
	Hash:=generatePwd(password)
	userModel,err:=dao.FindUserByNickName(nickname)
	if err!=nil{
		return err
	}

	if userModel.Password!=Hash{
		return errors.New("用户名或密码错误")
	}
	return


}


func generatePwd(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	PWD := hash.Sum(nil)
	return hex.EncodeToString(PWD)
}

func ValiatePwd(password string, hashpassword string) bool {
	return generatePwd(password) == hashpassword
}