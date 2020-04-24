package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v4"
	"github.com/joho/godotenv"
	_ "go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/xDarkicex/Hospice/helpers"
	"github.com/xDarkicex/Hospice/server"
	"github.com/xDarkicex/Hospice/terminal"
)

var myEnv map[string]string
var PORT string
var Email_ADDRESS string
var EMAIL_PASSWORD string
var GOOGLE_KEY string
var address string
var IP string
var PRODUCTION string
var DEVELOPMENT string
var ADDRESS string
var production bool
var development bool

func init() {

	path, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}
	myEnv, err = godotenv.Read(path + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Email_ADDRESS = myEnv["EMAIL_ADDRESS"]
	EMAIL_PASSWORD = myEnv["EMAIL_PASSWORD"]
	GOOGLE_KEY = myEnv["GOOGLE_KEY"]
	ADDRESS = myEnv["ADDRESS"]
	IP = myEnv["IP"]
	DEVELOPMENT = myEnv["DEVELOPMENT"]
	PRODUCTION = myEnv["PRODUCTION"]
	PORT = myEnv["PORT"]
	fmt.Println(myEnv)
	production, err := strconv.ParseBool(PRODUCTION)
	if err != nil {
		log.Print(err, production)
	}
	development, err := strconv.ParseBool(DEVELOPMENT)
	if err != nil {
		log.Print(err, development)
	}
}

// Port is env port system Port
func main() {

	s := server.NewRouter()
	var config = struct {
		Port              string
		IP                string
		Address           string
		ENV               *server.Env
		StartTime         time.Time
		Handler           *chi.Mux
		ReadTimeout       time.Duration
		ReadHeaderTimeout time.Duration
		WriteTimeout      time.Duration
		IdleTimeout       time.Duration
	}{
		ENV: &server.Env{
			Production:  production,
			Development: development,
		},
		Port:              PORT,
		IP:                IP,
		Address:           ADDRESS,
		Handler:           s,
		StartTime:         time.Now(),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	fmt.Println(fmt.Sprintf(helpers.Color.Coral("Server Architecture detected")+"=%s\n"+helpers.Color.Coral("Server CPU Count")+"=%s\n", helpers.Color.Blue(string(runtime.GOARCH)), helpers.Color.Blue(strconv.FormatInt(int64(runtime.NumCPU()), 10))))

	fmt.Println("[" + terminal.Colors[218]("200") + "]" + terminal.Colors[212]("Now Listening ") + helpers.Color.PinkBold("ON ") + helpers.Color.Blue(ADDRESS))
	err, srv := server.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}
	mux := srv.Handler
	c := context.Background()
	log.Fatalln(http.ListenAndServe(srv.Address, chi.ServerBaseContext(c, mux)))
	fmt.Println(c)

}
