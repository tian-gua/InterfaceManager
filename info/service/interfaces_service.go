package service

import (
	"github.com/aidonggua/growing/gorm"
	. "info/entity"
)

func FindInterfaces(moduleId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{ModuleId: moduleId}, ints)
	return *ints, err
}

func AddInterfaces(interf *Interfaces) error {
	return gorm.Save(interf)
}