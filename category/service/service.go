package service

import (
	"github.com/siliconvalley001/wen/category/dao"
	"context"
)

type Service_Category struct {
	ctx *context.Context
	dao *dao.Dao
}
