package service

import (
	. "info/entity"

	"github.com/aidonggua/growing/gorm"
)

func FindInterfaces(moduleId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{ModuleId: moduleId}, ints)
	return *ints, err
}

func AddInterfaces(interf *Interfaces) error {
	return gorm.Save(interf)
}

func FindInterfaceById(interfaceId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{Id: interfaceId}, ints)
	return *ints, err
}
