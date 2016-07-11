package controller

import (
	"net/http"
	. "index/service"
	"html/template"
	"github.com/aidonggua/growing/grouter"
	"common"
)

func index(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./template/view/index.html")
	if err != nil {
		panic(err)
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

type addParam struct {
	ModuleName string
	ProjectId  int
}

func add(rw http.ResponseWriter, req *http.Request, param addParam) {
	err := AddModule(param.ModuleName, param.ProjectId)
	if err != nil {
		rw.Write(common.GetCustomStatus("添加模块错误", 1001, nil).GetJson())
	} else {
		rw.Write(common.GetSuccessStatus().SetMsg("添加成功").GetJson())
	}
}

func init() {
	grouter.Route("/index", index)
	grouter.Route("/module/add", add)

}