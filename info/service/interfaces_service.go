package service

import (
	. "info/entity"

	"fmt"
	"strings"

	"github.com/aidonggua/growing/gorm"
)

func FindInterfaces(moduleId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{ModuleId: moduleId}, ints)
	return *ints, err
}

func AddInterfaces(interf *Interfaces, paramStr string) error {
	tx, err := gorm.Begin()
	if err != nil {
		//直接退出
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("事务回滚")
			tx.RollBack()
		}
	}()
	err = gorm.Save(interf, tx)
	if err != nil {
		//交给recover处理
		panic(err)
	}
	//获得刚插入的最新的id
	interfaceId := new(int)
	gorm.CustomQuery("select MAX(id) from interfaces", interfaceId, tx)
	//插入param
	paramStr = strings.TrimSpace(paramStr)
	if len(paramStr) > 0 {
		params := strings.Split(paramStr, ",")
		if len(params) > 0 {
			for _, v := range params {
				param := strings.Split(strings.TrimSpace(v), " ")
				if len(param) < 2 {
					fmt.Println("类型" + param[0] + "没有变量名!")
					continue
				}
				for i := 1; i < len(param); i++ {
					err := gorm.Save(&Params{InterfaceId: *interfaceId, Type: param[0], Name: param[i]}, tx)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
	tx.Commit()
	return nil
}

func FindInterfaceById(interfaceId int) ([]Interfaces, error) {
	ints := new([]Interfaces)
	err := gorm.Query(&Interfaces{Id: interfaceId}, ints)
	return *ints, err
}
