package server

import (

    "io"


    log2 "log"
	"net/http"
    "os"
    "strings"
	"time"
	log "github.com/labstack/gommon/log"
	"github.com/go-chi/chi/v4"
	"github.com/go-chi/chi/v4/middleware"

    "github.com/valyala/fasttemplate"

	"github.com/xDarkicex/cchha_new_server/app/controllers"
	"github.com/xDarkicex/cchha_new_server/helpers"
    "github.com/xDarkicex/cchha_new_server/terminal"
)

var handle = new(helpers.HandleLeveled)
func NewRouter() *chi.Mux {
    router := chi.NewRouter()
    router.Use(middleware.Timeout(60 * time.Second))

    // Routes
    // splash
    application := controllers.Application{}
    router.Get("/", application.Index)

    // home health
    homeHealth := controllers.HomeHealth{}
    router.Get("/home-health", homeHealth.Index)
    router.Get("/home-health/careers", homeHealth.Careers)
    router.Get("/home-health/careers.html", homeHealth.Careers)
    router.Get("/home-health/services", homeHealth.Services)
    router.Get("/home-health/services.html", homeHealth.Services)
    router.Get("/home-health/eligibility", homeHealth.Eligibility)
    router.Get("/home-health/eligibility.html", homeHealth.Eligibility)
    router.Get("/home-health/resources", homeHealth.Resources)
    router.Get("/home-health/resources.html", homeHealth.Resources)
    router.Get("/home-health/community", homeHealth.Community)
    router.Get("/home-health/community.html", homeHealth.Community)
    router.Get("/home-health/about", homeHealth.About)
    router.Get("/home-health/about.html", homeHealth.About)
    router.Get("/home-health/locations", homeHealth.Locations)
    router.Get("/home-health/locations.html", homeHealth.Locations)
    router.Get("/home-health/contact", homeHealth.Contact)
    router.Get("/home-health/contact.html", homeHealth.Contact)

    // hospice
    hospice := controllers.Hospice{}
    router.Get("/hospice", hospice.Index)
    router.Get("/hospice/careers", hospice.Careers)
    router.Get("/hospice/careers.html", hospice.Careers)
    router.Get("/hospice/services", hospice.Services)
    router.Get("/hospice/services.html", hospice.Services)
    router.Get("/hospice/eligibility", hospice.Eligibility)
    router.Get("/hospice/eligibility.html", hospice.Eligibility)
    router.Get("/hospice/resources", hospice.Resources)
    router.Get("/hospice/resources.html", hospice.Resources)
    router.Get("/hospice/community", hospice.Community)
    router.Get("/hospice/community.html", hospice.Community)
    router.Get("/hospice/about", hospice.About)
    router.Get("/hospice/about.html", hospice.About)
    router.Get("/hospice/locations", hospice.Locations)
    router.Get("/hospice/locations.html", hospice.Locations)
    router.Get("/hospice/contact", hospice.Contact)
    router.Get("/hospice/contact.html", hospice.Contact)



    // static assets
    router.Get("/static/{filepath}*", func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = chi.URLParam(r, "filepath")
        if strings.ContainsAny(r.URL.Path, "{}*") {
            log.Error("FileServer does not permit URL parameters.")
        }
        http.FileServer(http.Dir("public")).ServeHTTP(w, r)
    })

    return router
}

var Port = ":3000"

type server struct {
    Port      string
    Address   string
    ENV       e
    StartTime time.Time
    Router    *http.Server
}

type e struct {
    Production  bool
    Development bool
}
var Color = terminal.NewTerminalColor()
var Colors = terminal.Colors
var Mux *chi.Mux
var (
    Default *log2.Logger
    Trace   *log2.Logger
    Info    *log2.Logger
    Warn *log2.Logger
    Error   *log2.Logger
)
var (
    defaultPre = Color.Blue("[" ) + "{{level}}" + Color.Blue("]" ) +" :"
    infoPre = Color.PinkBold("[" ) + "{{level}}" + Color.PinkBold("]" ) +" :"
    tracePre = Color.Orange("[" ) + "{{level}}" + Color.Orange("]" ) +" :"
    errorPre = Color.GreenLight("[" ) + "{{level}}" + Color.GreenLight("]" ) +" :"
    warnPre = Color.Coral("[" ) + "{{level}}" + Color.Coral("]" ) +" :"
)

func Init(defaultHandle io.Writer, infoHandle io.Writer, warnHandle io.Writer, errorHandle io.Writer, traceHandle io.Writer) {

    Default = defaultLogger(defaultHandle)
    Error = errorLogger(errorHandle)
    Trace = traceLogger(traceHandle)
    Info = infoLogger(infoHandle)
    Warn = warnLogger(warnHandle)
}

func defaultLogger(dh io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(defaultPre, "{{","}}", map[string]interface{}{
        "level": Color.PinkLight("Default"),
    })
    defaultPRE := prefixed
    return log2.New(dh, defaultPRE, log2.Ltime|log2.LstdFlags)
}

func errorLogger(eh io.Writer) *log2.Logger {
	prefixed := fasttemplate.ExecuteString(errorPre, "{{","}}", map[string]interface{}{
		"level":  Color.Red("Error"),
	})
    coloredPRE := prefixed
	return log2.New(eh, coloredPRE, log2.Ldate|log2.Ltime|log2.Lshortfile)
}
func traceLogger(th io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(tracePre, "{{","}}", map[string]interface{}{
        "level": Color.Blue("Trace"),
    })
    coloredPRE := prefixed
    return log2.New(th, coloredPRE, log2.Ldate|log2.Ltime|log2.Lshortfile)
}
func infoLogger(ih io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(infoPre, "{{","}}", map[string]interface{}{
        "level": Color.GreenLight("Info"),
    })
    coloredPRE := prefixed
    return log2.New(ih, coloredPRE, log2.Ldate|log2.Ltime|log2.Lshortfile)
}
func warnLogger(wh io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(warnPre, "{{","}}", map[string]interface{}{
        "level": Color.Orange("Warning"),
    })
    return log2.New(wh, prefixed, log2.Ldate|log2.Ltime|log2.Lshortfile)
}
func init() {
    Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout,  os.Stderr)
    Default.Println("...")
    Info.Println(".....")
    Warn.Println("Loading....")

    Error.Println(Color.PinkBold("Done loading colors"))

    Mux = NewRouter()
    Error.Println(Color.Blue("Routes Registration complete"))

}

func NewServer(config interface{}) *http.Server {
    switch c := config.(type) {
        case map[string]interface{}:
            config, ok := config.(map[string]interface{})
            if !ok {
                Error.Println(Color.RedBlink("Config loaded incorrectly please look at configuration file for errors"))
                return nil
            }
            addr, ok := config["addr"].(string)
            if !ok {
                Error.Fatalln(addr)
            }
            handler, ok := config["handler"].(http.Handler)
            if !ok {
                Error.Fatalln(handler)
            }
            readTimeout, ok := config["readTimeout"].(time.Duration)
            if !ok {
                Error.Fatalln(readTimeout)
            }
            readHeaderTimeout, ok := config["readHeaderTimeout"].(time.Duration)
            if !ok {
                Error.Fatalln(readHeaderTimeout)
            }
            writeTimeout, ok := config["writeTimeout"].(time.Duration)
            if !ok {
                Error.Fatalln(writeTimeout)
            }
            ideaTimeout, ok := config["ideaTimeout"].(time.Duration)
            if !ok {
                Error.Fatalln(ideaTimeout)
            }
            return &http.Server{
                Addr:    addr,
                Handler: handler,
                ReadTimeout:      readTimeout,
                ReadHeaderTimeout: readHeaderTimeout,
                WriteTimeout:      writeTimeout,
                IdleTimeout:       ideaTimeout,
                ErrorLog:          &log2.Logger{},
            }
        case struct{
                Addr              string
                Handler           http.Handler
                ReadTimeout       time.Duration
                ReadHeaderTimeout time.Duration
                WriteTimeout      time.Duration
                IdleTimeout       time.Duration
            }:
            return &http.Server{
            Addr:    c.Addr,
            Handler: c.Handler,
            ReadTimeout:      c.ReadTimeout,
            ReadHeaderTimeout: c.ReadHeaderTimeout,
            WriteTimeout:      c.WriteTimeout,
            IdleTimeout:       c.IdleTimeout,
            ErrorLog:          &log2.Logger{},
        }
        case string:
            Error.Panicln("NO STRINGS!")
            return nil
            // make json string side
        case []byte:
            Error.Panicln("NO BYTES!")
            return nil
            // make json byte style
        default:
            Info.Println("Need to implement json and reverse json")
            Error.Panicln("FAILING AS HARD AS WE CAN")
            return nil
    }
    Error.Panic("NO ONE SHOULD BE HERE!")
    return nil
}
