package controllers

import (
	"net/http"

	//"github.com/golang/glog"

	//"html/template"

	"./helpers"
	//"github.com/haruyama/golang-goji-sample/models"
	"github.com/prashanthrv/goji-handlebars/system"
	"github.com/zenazn/goji/web"
)

type MainController struct {
	system.Controller
}

// Home page route
func (controller *MainController) Index(c web.C, r *http.Request) (string, int) {

	return system.RenderFile("default"), http.StatusOK
}