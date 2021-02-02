package service

import (
	"github.com/siliconvalley001/wen/product/dao"
	"github.com/siliconvalley001/wen/product/model"
)

func AddProduct(product *model.Product)(int64,error){
	return dao.CreateProduct(product)

}

func DelProduct(productid int64)error{
	return dao.DeleteProductById(productid)
}

func UpdateProduct(product *model.Product)error{
	return dao.UpdateProduct(product)
}

func FindProductById(productid int64)(*model.Product,error){
	return dao.FindProductById(productid)
}

func FindAllProduct()([]model.Product,error){
	return dao.FindAll()
}
