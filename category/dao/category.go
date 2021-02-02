package dao

import "github.com/siliconvalley001/wen/category/model"

func InitTable()error{
	return model.DB.CreateTable(&model.Category{}).Error
}
func CreateCategory(category *model.Category)(id int64,err error){
	return category.Id, model.DB.Create(category).Error
}
func UpdateCategory(category *model.Category)error{
	return model.DB.Model(category).Update(&category).Error

}
func DeleteCategoryById(id int64)error{
	return model.DB.Where(`id=?`, id).Delete(&model.Category{}).Error
}
func FindCategoryByName(name string) (user *model.Category, err error) {
	category:= &model.Category{}
	return user, model.DB.Where(`name= ?`, name).Find(category).Error
}
func FindCategoryByID(id int64)(cate *model.Category,err error){
	cate=&model.Category{}
	return cate, model.DB.First(cate, id).Find(cate).Error
}
func FindCategoryByLevel(level uint32)(allcate []model.Category,err error){
	return allcate,model.DB.Find(`level= ?`,level).Find(&allcate).Error

}
func FindCategoryByParent(parent int64)(allcate []model.Category,err error){
	return allcate,model.DB.Find(`parent= ?`,parent).Find(&allcate).Error

}

func FindAllCategory()(all []model.Category,err error){
	return all, model.DB.Find(&all).Error
}