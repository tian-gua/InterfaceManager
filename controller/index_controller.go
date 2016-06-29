package controller

import (
	"net/http"
	"service"
	"text/template"

	"github.com/aidonggua/growing/grouter"
	"fmt"
)

type IndexController struct {
	grouter.BaseController
}

func (this *IndexController) Get(rw http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./template/view/index.html")
	if err != nil {
		//panic(err)
	}
	modules, err := service.FindAllModule(0)
	if err != nil {
		panic(err)
	}
	fmt.Println(modules)
	data := template.FuncMap{"modules": modules}
	err = t.Execute(rw, data)
	if err != nil {
		panic(err)
	}
}
