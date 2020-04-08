package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	_ "github.com/labstack/gommon/log"

	"github.com/xDarkicex/Hospice/helpers"
	"github.com/xDarkicex/Hospice/server"
	"github.com/xDarkicex/Hospice/terminal"
)

func inti() {
	multiplex := io.MultiWriter(os.Stderr, os.Stdout)
	helpers.Init(multiplex, multiplex, multiplex, multiplex, multiplex)
}

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

	fmt.Println(fmt.Sprintf(helpers.Color.Coral("Server Architecture detected")+"=%s\n"+helpers.Color.Coral("Server CPU Count")+"=%s\n", helpers.Color.Blue(string(runtime.GOARCH)), helpers.Color.Blue(strconv.FormatInt(int64(runtime.NumCPU()), 10))))
	var address = config.Addr
	fmt.Println("[" + terminal.Colors[218]("200") + "]" + terminal.Colors[212]("Now Listening ") + helpers.Color.PinkBold("ON ") + helpers.Color.Blue(address))
	var srv = server.NewServer(config)
	log.Fatalln(srv.ListenAndServe().Error())

}
