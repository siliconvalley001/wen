package dao

import (
	"github.com/siliconvalley001/wen/user/model"
)

func  InitUserTable() error {
	return model.DB.CreateTable(&model.User{}).Error
}

func  FindUserByNickName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, model.DB.Where(`nick_name= ?`, name).Find(user).Error
}

func  FindUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	return user, model.DB.First(user, id).Find(user).Error

}

func  CreateUser(user *model.User) (id int64, err error) {


	return user.Id, model.DB.Create(user).Error
}

func  DelUserByID(id int64) error {
	return model.DB.Where(`id=?`, id).Delete(&model.User{}).Error

}

func  UpdateUser(user *model.User) error {
	return model.DB.Model(user).Update(&user).Error
}

func  FindAll() (userall []model.User, err error) {
	return userall, model.DB.Find(&userall).Error
}

