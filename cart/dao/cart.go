package dao

import (
	"github.com/siliconvalley001/wen/cart/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func  FindCartById(cart_id int64) (cart *model.Cart, err error) {

	cart = &model.Cart{}
	return cart, model.DB.First(cart, cart_id).Error
}

func  CreateCart(cart *model.Cart) (int64, error) {
	db :=  model.DB.FirstOrCreate(cart, model.Cart{ProductId: cart.ProductId, SizeId: cart.SizeId})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected <= 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.Id, nil
}
func  DelCartById(cart_id int64) error {
	return  model.DB.Where(`cart_id=?`, cart_id).Delete(&model.Cart{}).Error
}

func  UpdateCart(cart *model.Cart) error {
	return  model.DB.Model(cart).Update(cart).Error
}

func  FindAll(user_id int64) (all []model.Cart, err error) {
	return all,  model.DB.Where(`user_id=?`, user_id).Find(&all).Error
}

func  CleanCart(user_id int64) error {
	return  model.DB.Where(`user_id=?`, user_id).Delete(&model.Cart{}).Error
}

func  IncrNum(cartid int64, num int64) error {
	cart := &model.Cart{
		Id: cartid,
	}
	return  model.DB.Model(cart).UpdateColumn("num", gorm.Expr(`num=?`, num)).Error
}

func DecrNum(cartid int64, num int64) error {
	cart := &model.Cart{
		Id: cartid,
	}
	db:= model.DB.Model(cart).Where(`num> ?`,num).UpdateColumn("num", gorm.Expr(`num=?`, num))

	if db.Error!=nil{
		return db.Error
	}
	return nil


}
