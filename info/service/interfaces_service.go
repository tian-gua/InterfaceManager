package service

import (
	. "info/entity"
	"github.com/aidonggua/growing/gorm"
	"strconv"
)

func FindInterfaces(moduleId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{ModuleId: moduleId}, ints)
	return *ints, err
}

func AddInterfaces(interf *Interfaces) error {
	return gorm.Save(interf)
}

func FindInterfaceById(interfaceId int) (*Interfaces, error) {
	ints := new(Interfaces)
	sqlStr := "select * from interfaces where id = "
	sqlStr += strconv.Itoa(interfaceId)
	err := gorm.CustomQuery(sqlStr, ints)
	return ints, err
}
