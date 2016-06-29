package main

import (
	"controller"

	"github.com/aidonggua/growing/gorm"
	"github.com/aidonggua/growing/gorm/generate"
	"github.com/aidonggua/growing/grouter"
	"github.com/aidonggua/growing/grow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	start()
}

func start() {
	gorm.InitDB("mysql", "root:root@tcp(127.0.0.1:3306)/my")
	grouter.Route("/index", &controller.IndexController{})
	grouter.Route("/info", &controller.InfoController{})
	grow.Start(8080)
}

func generateCode() {
	generate.SetDBInfo("mysql", "root:root@tcp(127.0.0.1:3306)/my")
	generate.Generate("interfaces")
	generate.Generate("module")

}
