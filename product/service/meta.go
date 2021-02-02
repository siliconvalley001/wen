package service

import (
	"github.com/siliconvalley001/wen/product/dao"
	"github.com/siliconvalley001/wen/product/model"
)

func FindMenusByid(id int64)(menusList *model.MenusList,err error){
	return dao.FindMenusByid(id)
}

func FindMenusByname(name string)(menusList *model.MenusList,err error) {
	return  dao.FindMenusByname(name)

}

func InsertMeta(list model.MenusList)(int64,error) {
	return dao.InsertMeta(list)
}


func Select()([]model.List,error){
	return  dao.SelectAll()
}
