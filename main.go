package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi/v4"
	"github.com/joho/godotenv"
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

	PORT = myEnv["PORT"]
	fmt.Println(myEnv)

}

// SSL checks env variables
func SSL(myEnv map[string]string) (production, development bool) {
	DEVELOPMENT = myEnv["DEVELOPMENT"]
	PRODUCTION = myEnv["PRODUCTION"]
	production, err := strconv.ParseBool(PRODUCTION)
	if err != nil {
		log.Print(err, production)
	}
	development, err = strconv.ParseBool(DEVELOPMENT)
	if err != nil {
		log.Print(err, development)
	}
	return production, development
}

// Port is env port system Port
func main() {
	served := server.NewRouter()
	production, development := SSL(myEnv)
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
		Handler:           served,
		StartTime:         time.Now(),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	fmt.Println(fmt.Sprintf(helpers.Color.Coral("Server Architecture detected")+"=%s\n"+helpers.Color.Coral("Server CPU Count")+"=%s\n", helpers.Color.Blue(string(runtime.GOARCH)), helpers.Color.Blue(strconv.FormatInt(int64(runtime.NumCPU()), 10))))

	fmt.Println("[" + terminal.Colors[218]("200") + "]" + terminal.Colors[212]("Now Listening ") + helpers.Color.PinkBold("ON ") + helpers.Color.Blue(ADDRESS))

	if config.ENV.Production {
		fmt.Println("...")
	}
	// err = http.ListenAndServe(srv.Address, chi.ServerBaseContext(c, mux))
	// if err != nil {
	// 	fmt.Println(c)
	// }

}
