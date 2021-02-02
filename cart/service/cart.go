package service

import (
	"github.com/siliconvalley001/wen/cart/dao"
	"github.com/siliconvalley001/wen/cart/model"
)

func AddCart(cart *model.Cart)(int64,error){
	return dao.CreateCart(cart)

}

func DelCart(cartid int64)error{
	return dao.DelCartById(cartid)
}


func UpdateCart(cart *model.Cart)error{
	return dao.UpdateCart(cart)
}

func FindCartByID(cartid int64)(*model.Cart,error){
	return dao.FindCartById(cartid)
}

func FindCartAll(userid int64)([]model.Cart,error)  {
	return dao.FindAll(userid)
}


func CleanCart(userid int64)error{
	return dao.CleanCart(userid)
}


func DecrNum(cartid int64,num int64)error{
	return  dao.DecrNum(cartid,num)
}


func IncrNum(cartid int64,num int64)error{
	return dao.IncrNum(cartid,num)
}