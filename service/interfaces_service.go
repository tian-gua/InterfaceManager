package service

import (
	. "github.com/aidonggua/InterfaceManager/domain"
	"github.com/aidonggua/InterfaceManager/dao"
)

//根据模块id查询模块
func FindInterfaces(moduleId int) ([]Interfaces, error) {
	return dao.SelectInterfacesByModuleId(moduleId)
}

//添加接口
func AddInterfaces(interf *Interfaces) error {
	return dao.InsertInterfaces(interf)
}
//根据id查询指定接口
func FindInterfaceById(id int) (*Interfaces, error) {
	return dao.SelectInterfacesById(id);
}
