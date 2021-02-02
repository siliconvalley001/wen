package handler

import (
	"context"
	"github.com/siliconvalley001/wen/product/common"
	ex "github.com/siliconvalley001/wen/product/proto"
	"github.com/siliconvalley001/wen/product/service"
)

type Meta struct {
}

func (m *Meta) GetMetaById(ctx context.Context, resquest *ex.MetaByIdRequest, response *ex.MetaResponse) error {
	_, err := service.FindMenusByid(resquest.Id)
	if err != nil {
		return err
	}
	response.Msg = "获取菜单成功"
	response.Status = 200
	return nil
}
func (m *Meta) GetMetaByName(ctx context.Context, resquest *ex.MetaByNameRequest, response *ex.MetaResponse) error {
	_, err := service.FindMenusByname(resquest.Name)
	if err != nil {
		return err
	}
	response.Msg = "获取菜单成功"
	response.Status = 200
	return nil
}
func (m *Meta) GetMetaAll(ctx context.Context, resquest *ex.MetaAllRequest, response *ex.MetaAllResponse) error {

	all, err := service.FindAllProduct()
	if err != nil {
		return err
	}
	for _,value:=range all{
		metaInfo:=&ex.MetaInfo{}
		err:=common.Swap(value,metaInfo)
		if err!=nil{
			return err
		}
		response.Info=append(response.Info,metaInfo)
	}
	return nil
}
