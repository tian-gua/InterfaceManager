package service

import (
	"github.com/aidonggua/growing/gorm"
	"entity"
)

//获得所有模块
func FindAllModule(projectId int) ([]entity.Module, error) {

	modules := new([]entity.Module)
	err := gorm.QueryAll(modules)
	return *modules, err

}
