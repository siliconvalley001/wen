package handler

import (
	"github.com/siliconvalley001/wen/cart/common"
	"github.com/siliconvalley001/wen/cart/model"
	"github.com/siliconvalley001/wen/cart/service"
	"context"
	ex "github.com/siliconvalley001/wen/cart/proto"

)

type Cart struct{

}


func (c *Cart)	AddCart(ctx context.Context, resquest *ex.CartInfo, response *ex.ResponseAdd) error{
	cart:=&model.Cart{}
	//common.Swap(resquest,cart)
	common.Swap(resquest,cart)

	if _,err:=service.AddCart(cart);err!=nil{
		return err
	}
	return nil

}

func (c *Cart)CleanCart(ctx context.Context,resquest *ex.Clean, response *ex.Response) error{
	if err:=service.CleanCart(resquest.UserId);err!=nil{
		return err
	}
	response.Meg="购物车清空成功"
	return nil

}
func (c *Cart)Incr(ctx context.Context,resquest *ex.Item, response *ex.Response) error{
	if err:=service.IncrNum(resquest.Id,resquest.ChangeNum);err!=nil{
		return err
	}
	response.Meg="购物车新增成功"
	return nil
}
func (c *Cart)Decr(ctx context.Context, resquest *ex.Item,response *ex.Response) error{
	if err:=service.DecrNum(resquest.Id,resquest.ChangeNum);err!=nil{
		return err
	}
	response.Meg="购物车减少成功"
	return nil

}
func (c *Cart)DeleteItemByID(ctx context.Context,resquest *ex.CartID,response *ex.Response) error{
	if err:=service.DelCart(resquest.Id);err!=nil{
		return err
	}
	response.Meg="购物车删除成功"
	return nil
}
func (c *Cart)GetAll(ctx context.Context, resquest *ex.CartFindAll,response *ex.CartAll) error{
	all,err:=service.FindCartAll(resquest.UserId)
	if err!=nil{
		return err
	}
	for _,v:=range all{
		cart:=&ex.CartInfo{}
		if err:=common.Swap(v,cart);err!=nil{
			return err
		}
		response.CartInfo=append(response.CartInfo,cart)



	}
	return nil

}



