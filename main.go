package main

import (
	_ "controller"
	"github.com/aidonggua/growing/gorm"
	"github.com/aidonggua/growing/gorm/generate"
	"github.com/aidonggua/growing/grow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	start()
	//	generateCode()
}

func start() {
	gorm.InitDB("mysql", "root:root@tcp(127.0.0.1:3306)/my")
	grow.Start(8080)
}

func generateCode() {
	generate.SetDBInfo("mysql", "root:root@tcp(127.0.0.1:3306)/my")
	generate.Generate("interfaces")
	generate.Generate("module")
	generate.Generate("params")
}
