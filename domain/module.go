package domain

import "github.com/aidonggua/growing/gorm"

type Module struct {
	Id        int    //编号
	Name      string //模块名
	ProjectId int    //项目id
}


//获得模块下的所有接口
//在template里面调用
func (m *Module) GetInterfaces() []Interfaces {
	itfcs := new([]Interfaces)
	err := gorm.Query(&Interfaces{ModuleId: m.Id}, itfcs)
	if err != nil {
		panic(err)
	}
	return *itfcs
}