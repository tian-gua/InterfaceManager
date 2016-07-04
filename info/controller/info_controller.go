package controller

import (
	"html/template"
	"net/http"
	"github.com/aidonggua/growing/grouter"
)

func info(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./template/view/info.html")
	if err != nil {
		panic(err)
	}

	t.Execute(rw, nil)
}

func init() {
	grouter.Route("/info", info)
}