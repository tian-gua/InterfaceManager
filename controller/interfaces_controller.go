package controller

import (
	"net/http"
	"html/template"
	"service"
	"common"
	"strconv"
	. "domain"
	"github.com/aidonggua/growing/grouter"
	"github.com/aidonggua/growing/gorm"
)

const md string = `

### 描述：
	描述一下这个接口

------------

###请求地址：
	Http://localhost:8080/login

------------

###请求类型：
	Get

------------

###请求参数
| 参数名  | 参数类型  | 说明 | 是否必选 |
| ------------ | ------------ | ------------ | ------------ |
|	userName	|	String	|	用户名	|	是	|
|	password	|	String	|	密码	|	是	|
|	mobile	|	String	|	手机号	|	否	|

------------

###返回值
| 参数名  | 参数类型  | 说明 | 是否必选 |
| ------------ | ------------ | ------------ | ------------ |
|	code	|	String	|	返回码	|	是	|
|	msg	|	String	|	接口信息	|	是	|
|	result	|	String	|	结果集	|	否	|


------------

###返回实例
	{
	"statu": "success",
	"code": "000",
	"result": null
	}

`

type infoParam struct {
	Id   int
	Type int
}

func info(rw http.ResponseWriter, req *http.Request, param infoParam) {
	t, err := template.ParseFiles("./template/view/interfaces.html")
	if err != nil {
		panic(err)
	}
	interf, err := service.FindInterfaceById(param.Id)
	if err != nil {
		panic(err)
	}
	t.Execute(rw, interf)
}

func markdown(rw http.ResponseWriter, req *http.Request, param infoParam) {
	t, err := template.ParseFiles("./template/view/markdown.html")
	if err != nil {
		panic(err)
	}
	var itfc *Interfaces = nil
	if param.Type == 0 {
		itfc = &Interfaces{Id:0, ModuleId: param.Id, Html: md, Name:""}
	} else {
		itfc, err = service.FindInterfaceById(param.Id)
		if err != nil {
			panic(err)
		}
	}
	t.Execute(rw, itfc)
}

//保存接口
func addInterfaces(rw http.ResponseWriter, req *http.Request, intelf Interfaces) {
	if len(intelf.Name) == 0 || intelf.ModuleId == 0 || len(intelf.Html) == 0 {
		rw.Write(common.GetCustomStatus("参数不正确!", 4000, nil).GetJson())
		return
	}
	err := service.AddInterfaces(&intelf)
	if err != nil {
		if err == gorm.ZRC {
			rw.Write(common.GetCustomStatus("没有做任何修改!", 200, nil).GetJson())
		} else {
			rw.Write(common.GetCustomStatus("保存接口错误!", 1003, nil).GetJson())
		}
	} else {
		rw.Write(common.GetSuccessStatus().SetMsg("保存接口成功!").GetJson())
	}

}

//查找指定模块下的接口
func findInterfaces(rw http.ResponseWriter, req *http.Request) {
	moduleId, err := strconv.Atoi(req.FormValue("moduleId"))
	if err != nil {
		panic(err)
	}
	ints, err := service.FindInterfaces(moduleId)
	if err != nil {
		rw.Write(common.GetCustomStatus("获取接口错误!", 1002, nil).GetJson())
	} else {
		rw.Write(common.GetSuccessStatus().SetMsg("获取接口成功!").SetResult(ints).GetJson())
	}
}

func init() {
	grouter.Route("/info", info)
	grouter.Route("/interfaces/find", findInterfaces)
	grouter.Route("/interfaces/add", addInterfaces)
	grouter.Route("/markdown", markdown)
}
