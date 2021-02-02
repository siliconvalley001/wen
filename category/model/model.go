package model

import (
	"github.com/siliconvalley001/project1/category/common"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)
var (
	DB *gorm.DB
	err error
)


func init(){
	if _,err:=common.InitSetting();err!=nil{
		panic(err)
	}
	fmt.Println("ll",		common.Con.Mys.Name)

	DB,err=gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		common.Con.Mys.Name,
		common.Con.Mys.PassWord,
		common.Con.Mys.Host,
		common.Con.Mys.Port,
		common.Con.Mys.DBName,
	))
	if err!=nil{
		panic(err)
	}
	DB.SingularTable(true)

	DB.DB().SetMaxOpenConns(common.Con.Mys.MaxOpenConns)

	DB.DB().SetMaxIdleConns(common.Con.Mys.MaxIdleConns)

	DB.AutoMigrate(&Category{})


}
