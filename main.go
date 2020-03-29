package main

import (
	"log"
	"net/http"


	"github.com/xDarkicex/cchha_new_server/server"

)


func main() {
	// var s = server.NewServer("127.0.0.1", ":3000")
	var routes = server.NewRouter()
	ser := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}
	log.Fatalln(ser.ListenAndServe())

}
