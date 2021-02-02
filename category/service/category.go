package service

import (
	"github.com/siliconvalley001/wen/category/dao"
	"github.com/siliconvalley001/wen/category/model"
)

func AddCategory(cate *model.Category)(int64,error){
	return dao.CreateCategory(cate)
}


func DelCategory(id int64)error{
	return dao.DeleteCategoryById(id)
}

func UpdateCategoryByID(id int64)error{
	return dao.DeleteCategoryById(id)
}
func FindCategoryByName(catename string)(cate *model.Category,err error){
	return dao.FindCategoryByName(catename)
}
func FindCategoryById(id int64)(cate *model.Category,err error){
	return dao.FindCategoryByID(id)
}
func FindCategoryByLevel(level uint32)(all []model.Category,err error){
	return dao.FindCategoryByLevel(level)
}
func FindCategoryByParent(parent int64)(all []model.Category,err error){
	return dao.FindCategoryByParent(parent)
}
func FindCategoryAll()(all []model.Category,err error){
	return dao.FindAllCategory()
}
