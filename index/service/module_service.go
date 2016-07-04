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

//添加模块
func AddModule(moduleName string, projectId int) error {
	module := &Module{Name:moduleName, ProjectId:projectId, }
	return gorm.Save(module)
}