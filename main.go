package main

import (
	"flag"
	"net/http"
	"fmt"
	
	"github.com/golang/glog"
	//"github.com/gorilla/context"

	//"github.com/prashanthrv/goji-handlebars/controllers"
	"github.com/prashanthrv/goji-handlebars/system"

	"github.com/zenazn/goji"
	//"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

func root(c web.C, w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hey")
}

func main() {
	filename := flag.String("config","config.toml","Path to configuration file")
	
	flag.Parse()
	defer glog.Flush()

	var application = &system.Application{}
	application.Init(filename)
	
	static := web.New()
	publicPath := application.Config.Get("general.public_path").(string)
	static.Get("/assets/*",http.StripPrefix("/assets/", http.FileServer(http.Dir(publicPath))))
	http.Handle("/assets/",static)

	goji.Get("/",root)
	
	goji.Serve()
}

