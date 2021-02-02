package handler

import (
	"github.com/siliconvalley001/wen/product/common"
	"github.com/siliconvalley001/wen/product/model"
	"github.com/siliconvalley001/wen/product/service"
	ex "github.com/siliconvalley001/wen/product/proto"
	"context"
)

type Product struct{
}



func (p *Product)AddProduct(ctx context.Context, resquest *ex.ProductInfo, response *ex.ResponseProduct) error{
	product:=&model.Product{}
	if err:=common.Swap(resquest,product);err!=nil{
		return err
	}
	productid,err:=service.AddProduct(product)
	if err!=nil{
		return err
	}
	response.ProductId=productid
	return nil
}
func (p *Product)FindProductByID(ctx context.Context,resquest  *ex.RequestID, response *ex.ProductInfo) error{
	product,err:=service.FindProductById(resquest.ProductId)
	if err!=nil{
		return err
	}
	if err:=common.Swap(resquest,product);err!=nil{
		return err
	}

	return nil
}
func (p *Product)UpdateProduct(ctx context.Context, resquest *ex.ProductInfo, response *ex.Response) error{
	product:=&model.Product{}
	if err:=common.Swap(resquest,product);err!=nil{
		return err
	}
	if err:=service.UpdateProduct(product);err!=nil{
		return err
	}
	response.Msg="更新成功"
	return nil

}
func (p *Product)DeleteProductByID(ctx context.Context,resquest *ex.RequestID, response *ex.Response) error{
	if err:=service.DelProduct(resquest.ProductId);err!=nil{
		return err
	}
	response.Msg="删除成功"
	return nil

}
func (p *Product)FindAllProduct(ctx context.Context, resquest *ex.RequestAll,response *ex.AllProduct) error{
	productall,err:=service.FindAllProduct()
	if err!=nil{
		return err
	}
	for _,value:=range productall{
		productinfo:=&ex.ProductInfo{}
		err:=common.Swap(value,productinfo)
		if err!=nil{
			return err
		}
		response.ProductInfo=append(response.ProductInfo,productinfo)
	}
	return nil
}
