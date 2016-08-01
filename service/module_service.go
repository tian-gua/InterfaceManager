package service

import (
	. "github.com/aidonggua/InterfaceManager/domain"
	"github.com/aidonggua/InterfaceManager/dao"
)

//获得所有模块
func FindAllModule(projectId int) ([]Module, error) {
	return dao.SelectAllModule(projectId)
}

//添加模块
func AddModule(moduleName string, projectId int) error {
	module := &Module{Name:moduleName, ProjectId:projectId, }
	return dao.InsertModule(module)
}
