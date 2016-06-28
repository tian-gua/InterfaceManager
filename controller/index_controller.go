package controller

import (
	"text/template"
	"net/http"

	"github.com/aidonggua/growing/grouter"
)

type IndexController struct {
	grouter.BaseController
}

func (this *IndexController) Get(rw http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./template/view/index.html")
	if err != nil {
		panic(err)
	}

	indexs := [5]string{"1", "1", "1", "1", "1"}
	data := template.FuncMap{"Index":indexs}
	err = t.Execute(rw, data)
	if err != nil {
		panic(err)
	}
}
