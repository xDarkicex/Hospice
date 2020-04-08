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
	"github.com/go-chi/chi/v4"
	"github.com/joho/godotenv"
	_ "github.com/labstack/gommon/log"

	"github.com/xDarkicex/Hospice/helpers"
	"github.com/xDarkicex/Hospice/server"
	"github.com/xDarkicex/Hospice/terminal"
)

func inti() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	multiplex := io.MultiWriter(os.Stderr, os.Stdout)
	helpers.Init(multiplex, multiplex, multiplex, multiplex, multiplex)
}
var Email_ADDRESS = os.Getenv("EMAIL_ADDRESS")
var EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
var GOOGLE_KEY = os.Getenv("GOOGLE_KEY")
var address = os.Getenv("ADDRESS")
var IP = os.Getenv("IP")

// Port is env port system Port
var Port = os.Getenv("PORT")
func main() {
	production, err := strconv.ParseBool(os.Getenv("Production"))
	if err != nil {
		log.Print(err)
	}
	development, err := strconv.ParseBool(os.Getenv("Development"))
	if err != nil {
		log.Print(err)
	}

	s := server.NewRouter()
	var config = struct {
			Port      string
			IP string
			Address   string
			ENV       *server.Env
			StartTime time.Time
			Handler *chi.Mux
			ReadTimeout       time.Duration
			ReadHeaderTimeout time.Duration
			WriteTimeout      time.Duration
			IdleTimeout       time.Duration
	}{
		ENV: &server.Env{
			Production: production,
			Development: development,
		},
		Port: Port,
		IP: IP,
		Address:              address + ":"+ Port,
		Handler:           s,
		StartTime: time.Now(),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	fmt.Println(fmt.Sprintf(helpers.Color.Coral("Server Architecture detected")+"=%s\n"+helpers.Color.Coral("Server CPU Count")+"=%s\n", helpers.Color.Blue(string(runtime.GOARCH)), helpers.Color.Blue(strconv.FormatInt(int64(runtime.NumCPU()), 10))))

	fmt.Println("[" + terminal.Colors[218]("200") + "]" + terminal.Colors[212]("Now Listening ") + helpers.Color.PinkBold("ON ") + helpers.Color.Blue(address))
	err, srv := server.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}
	mux := srv.Handler

	log.Fatalln(http.ListenAndServe(srv.Address, mux))

}
