package main

import (
    "fmt"
    "net/http"
    "runtime"
    "strconv"
    "time"
    _ "github.com/labstack/gommon/log"
    "github.com/xDarkicex/cchha_new_server/server"
)

func main() {
    var config = struct {
        Addr              string
        Handler           http.Handler
        ReadTimeout       time.Duration
        ReadHeaderTimeout time.Duration
        WriteTimeout      time.Duration
        IdleTimeout       time.Duration
    }{
        Addr:              "127.0.0.1:3000",
        Handler:           server.NewRouter(),
        ReadTimeout:       15 * time.Second,
        ReadHeaderTimeout: 5 * time.Second,
        WriteTimeout:      10 * time.Second,
        IdleTimeout:       20 * time.Second,
    }

    fmt.Println(server.Color.Coral("server successfully configured"), server.Color.Blue("!"))
    fmt.Println(fmt.Sprintf(server.Color.Coral("Server Architecture detected")+"=%s\n"+ server.Color.Coral("Server CPU Count")+"=%s\n", server.Color.Blue(string(runtime.GOARCH)), server.Color.Blue(strconv.FormatInt(int64(runtime.NumCPU()), 10))))
    var address = config.Addr
    fmt.Println("["+server.Colors[218]("200")+"]" + server.Colors[212]("Now Listening ") + server.Color.PinkBold("ON ") + server.Color.Blue(address))
    srv := server.NewServer(config)
    server.Error.Fatalln(server.Color.RedBlink(srv.ListenAndServe().Error()))


}
