package dao

import (
	. "domain"
	"github.com/aidonggua/growing/gorm"
)

//通过项目名查询所有模块
func SelectAllModule(id int) ([]Module, error) {
	modules := new([]Module)
	err := gorm.QueryAll(modules)
	return *modules, err
}

//插入module记录
func InsertModule(module *Module) error {
	return gorm.Save(module)
}