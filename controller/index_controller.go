package controller

import (
	"html/template"
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
	t.Execute(rw, nil)
}
