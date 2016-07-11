package entity

import . "github.com/aidonggua/InterfaceManager/info/entity"
import "github.com/aidonggua/InterfaceManager/info/service"

type Module struct {
	Id        int
	Name      string
	ProjectId int
}

func (m *Module) GetInterfaces() []Interfaces {
	ints, err := service.FindInterfaces(m.Id)
	if err != nil {
		panic(err)
	}
	return ints
}