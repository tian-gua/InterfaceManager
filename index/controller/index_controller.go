package controller

import (
	"net/http"
	. "index/service"
	"html/template"
	"github.com/aidonggua/growing/grouter"
	"fmt"
)

func index(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./template/view/index.html")
	if err != nil {
		//panic(err)
	}
	modules, err := FindAllModule(0)
	if err != nil {
		panic(err)
	}
	data := template.FuncMap{"modules": modules}
	err = t.Execute(rw, data)
	if err != nil {
		panic(err)
	}
}

func addModule(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("添加模块!")
}

func init() {
	grouter.Route("/index", index)
	grouter.Route("/addModule", addModule)
}