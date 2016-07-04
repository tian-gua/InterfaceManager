package controller

import (
	"html/template"
	"net/http"
	"github.com/aidonggua/growing/grouter"
	. "info/service"
	. "info/entity"
	"strconv"
	"common"
)

func info(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./template/view/info.html")
	if err != nil {
		panic(err)
	}

	t.Execute(rw, nil)
}

func add(rw http.ResponseWriter, req *http.Request) {
	moduleId, err := strconv.Atoi(req.FormValue("moduleId"))
	if err != nil {
		panic(err)
	}
	name := req.FormValue("name")
	description := req.FormValue("description")
	method, err := strconv.Atoi(req.FormValue("method"))
	if err != nil {
		panic(err)
	}
	int := &Interfaces{Name:name, ModuleId:moduleId, Description:description, Method:method}
	err = AddInterfaces(int)
	if err != nil {
		rw.Write(common.GetCustomStatus("添加接口错误!", 1003, nil).GetJson())
	} else {
		rw.Write(common.GetSuccessStatus().SetMsg("添加接口成功!").GetJson())
	}

}

func findInterfaces(rw http.ResponseWriter, req *http.Request) {
	moduleId, err := strconv.Atoi(req.FormValue("moduleId"))
	if err != nil {
		panic(err)
	}
	ints, err := FindInterfaces(moduleId)
	if err != nil {
		rw.Write(common.GetCustomStatus("获取接口错误!", 1002, nil).GetJson())
	} else {
		rw.Write(common.GetSuccessStatus().SetMsg("获取接口成功!").SetResult(ints).GetJson())
	}
}
func init() {
	grouter.Route("/info", info)
	grouter.Route("/interfaces/find", findInterfaces)
	grouter.Route("/interfaces/add", add)
}