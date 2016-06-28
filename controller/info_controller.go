package controller

import (
	"html/template"
	"net/http"

	"github.com/aidonggua/growing/grouter"
)

type InfoController struct {
	grouter.BaseController
}

func (this *InfoController) Get(rw http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./template/view/info.html")
	if err != nil {
		panic(err)
	}

	t.Execute(rw, nil)

}
