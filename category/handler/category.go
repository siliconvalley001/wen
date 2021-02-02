package handler

import (
	"github.com/siliconvalley001/wen/category/common"
	"github.com/siliconvalley001/wen/category/model"
	ex "github.com/siliconvalley001/wen/category/proto"
	"github.com/siliconvalley001/wen/category/service"
	"context"
	"log"
)

type Category struct {
}

func (h* Category)  CreateCategory(context context.Context, req *ex.CategoryRequest, resp *ex.CreateCategoryResponse) error {
	cate := &model.Category{}
	err := common.Swap(req, cate)
	if err != nil {
		return err
	}
	id, err := service.AddCategory(cate)
	if err != nil {
		return err
	}
	resp.CategoryId = id
	resp.Msg = "分类添加成功"
	return nil

}

func  (h* Category)UpdateCategory(con context.Context, req *ex.CategoryRequest,resp  *ex.UpdateCategoryResponse) error {
	cate := &model.Category{}
	err := common.Swap(req, cate)
	if err != nil {
		return err
	}
	err= service.UpdateCategoryByID(cate.Id)

	if err!=nil{
		return err
	}

	resp.Msg="分类更新成功"
	return nil
}
 func (h* Category)DeleteCategory(con context.Context,req *ex.DeleteCategoryResquest,resp *ex.DeleteCategoryResponse) error {
 	cate:=&model.Category{}
	 err := common.Swap(req, cate)
	 if err != nil {
		 return err
	 }
	err=service.DelCategory(req.CategoryId)
	if err!=nil{
		return err
	}
	resp.Msg="分类删除成功"
	return nil
 }
func (h* Category)FindCategoryByName(con context.Context, resq *ex.CategoryByNameResquest, resp *ex.CategoryResponse) error{

	cate,err:=service.FindCategoryByName(resq.CategoryName)
	if err!=nil{
		return err
	}
	return common.Swap(cate,resp)

}
func (h* Category)FindCategoryById(ctx context.Context, resquest *ex.CategoryByIdResquest, response *ex.CategoryResponse) error{
	cate, err := service.FindCategoryById(resquest.CategoryId)
	if err!=nil{
		return err
	}
	return common.Swap(cate,response)
}

func (h* Category)FindCategoryByLevel(ctx context.Context, resquest *ex.CategoryByLevelResquest,response *ex.FindAllResponse) error{
	cate,err:=service.FindCategoryByLevel(resquest.CategoryLevel)
	if err!=nil{
		return err
	}
	cateToResponse(cate,response)
	return nil
}

func cateToResponse(cate []model.Category,response *ex.FindAllResponse){
	for _,value:=range cate{
		c:=&ex.CategoryResponse{}
		err:=common.Swap(value,c)
		if err!=nil{
			log.Fatal(err)
			break
		}
		response.All=append(response.All,c)
	}

}

func  (h* Category)FindCategoryByParent(ctx context.Context, resquest *ex.CategoryByParentResquest, response *ex.FindAllResponse) error {
	all,err:=service.FindCategoryByParent(resquest.ParentId)
	if err!=nil{
		return err
	}
	cateToResponse(all,response)
	return nil
	
}

func (h* Category)FindAllCategory(ctx context.Context,  resquest*ex.FindAllResquest, response *ex.FindAllResponse) error{
	all,err:=service.FindCategoryAll()
	if err!=nil{
		return err
	}
	cateToResponse(all,response)
	return nil
}
