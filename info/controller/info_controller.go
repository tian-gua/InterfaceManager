package controller

import (
	"common"
	"html/template"
	. "info/entity"
	. "info/service"
	"net/http"
	"strconv"

	"github.com/aidonggua/growing/grouter"
)

func info(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./template/view/info.html")
	if err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		panic(err)
	}
	interf, err := FindInterfaceById(id)
	if err != nil {
		panic(err)
	}
	t.Execute(rw, interf[0])
}

func add(rw http.ResponseWriter, req *http.Request) {
	moduleId, err := strconv.Atoi(req.FormValue("moduleId"))
	if err != nil {
		panic(err)
	}
	name := req.FormValue("name")
	description := req.FormValue("description")
	url := req.FormValue("url")
	method, err := strconv.Atoi(req.FormValue("method"))
	if err != nil {
		panic(err)
	}
	int := &Interfaces{Name: name, ModuleId: moduleId, Description: description, Method: method, Url:url}
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
