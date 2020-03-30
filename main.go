package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v4"

    "github.com/xDarkicex/cchha_new_server/helpers"
    "github.com/xDarkicex/cchha_new_server/server"
)

func main() {
    h, mux := helpers.TTLERRORX()
    // var routes = server.NewRouter()
    ser := &http.Server{
        Addr:    ":3000",
        Handler: mux,

    }
    log.Fatalln(ser.ListenAndServe())

}
