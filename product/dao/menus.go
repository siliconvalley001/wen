package dao

import (
	"github.com/pkg/errors"
	"github.com/siliconvalley001/wen/product/model"
	"log"
)

func FindMenusByid(id int64)(menusList *model.MenusList,err error){
	menusList=&model.MenusList{}
	return menusList,model.DB.Where(`id= ?`, id).Find(menusList).Error
}



func FindMenusByname(name string)(menusList *model.MenusList,err error){
	menusList=&model.MenusList{}
	return menusList,model.DB.Where(`name= ?`, name).Find(menusList).Error
}

func InsertMeta(list model.MenusList)(int64,error)  {

	db := model.DB.FirstOrCreate(list, model.MenusList{Id: list.Id, Authname: list.Authname,})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected <= 0 {
		return 0, errors.New("InsertMeta")
	}
	return list.Id,nil
}
func SelectAll()( []model.List, error){
	all:=[]model.MenusList{}

	if model.DB.Find(&all).Error!=nil{
		return nil,model.DB.Find(&all).Error
	}
	total:=make([]model.List,len(all))
	for i:=0;i<len(all);i++{
		child,err:=findChildrenByid(all[i].ChildrenId)
		if err!=nil{
			log.Println("数据库操作失败")
			return nil, err
		}
		total[i].Id=all[i].Id
		total[i].Authname=all[i].Authname
		total[i].Path=all[i].Authname
		total[i].Children=append(total[i].Children,child...)
	}
	return total,nil
}

func findChildrenByid(id int64)(child []model.ChildrenList,err error){
	return child,model.DB.Where(`children_id= ?`,id).Find(&child).Error
}

