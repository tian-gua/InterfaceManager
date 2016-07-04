package entity

import . "info/entity"
import "info/service"

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