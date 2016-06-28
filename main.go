package main

import (
	"controller"
	"github.com/aidonggua/growing/grouter"
	"github.com/aidonggua/growing/grow"
)

func main() {

	grouter.Route("/index", &controller.IndexController{})
	grouter.Route("/info", &controller.InfoController{})
	grow.Start(8080)
}

