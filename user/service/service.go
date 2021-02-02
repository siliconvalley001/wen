package service

import (
	"context"
	"github.com/siliconvalley001/wen/user/dao"
)

type Service_User struct {
	ctx *context.Context
	dao *dao.Dao
}
