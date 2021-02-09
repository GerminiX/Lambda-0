package main

import (
	"log"
	"net/http"

	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/lambda-0/base-common/controllers/routers"
	"github.com/urfave/negroni"
)

func main() {
	base_common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr: base_common.AppConf.Server,
		Handler: n,
	}

	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

