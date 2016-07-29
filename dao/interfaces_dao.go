package dao

import (
	. "domain"
	"github.com/aidonggua/growing/gorm"
	"strconv"
)

//根据模块id查询接口列表
func SelectInterfacesByModuleId(moduleId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{ModuleId: moduleId}, ints)
	return *ints, err
}

//根据接口id查询接口
func SelectInterfacesById(id int) (*Interfaces, error) {
	ints := new(Interfaces)
	sqlStr := "select * from interfaces where id = "
	sqlStr += strconv.Itoa(id)
	err := gorm.CustomQuery(sqlStr, ints)
	return ints, err
}

//插入接口记录
func InsertInterfaces(interf *Interfaces) error {
	return gorm.Save(interf)
}
