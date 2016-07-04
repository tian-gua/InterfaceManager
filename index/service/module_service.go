package service

import (
	"github.com/aidonggua/growing/gorm"
	. "index/entity"
)

//获得所有模块
func FindAllModule(projectId int) ([]Module, error) {

	modules := new([]Module)
	err := gorm.QueryAll(modules)
	return *modules, err

}
