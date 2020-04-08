package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v4"
	"github.com/go-chi/chi/v4/middleware"
	"github.com/google/uuid"
	log "github.com/labstack/gommon/log"
	"github.com/valyala/fasttemplate"

	"github.com/xDarkicex/Hospice/app/controllers"
	"github.com/xDarkicex/Hospice/helpers"
	"github.com/xDarkicex/Hospice/terminal"
)

var Error = helpers.Error

func init() {

}

func SiteByHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxRaw := r.Context()
		ctx := context.WithValue(ctxRaw, "x-author", "https://bitdev.io")
		ctx = context.WithValue(ctx, "siteby", "bitdev")
		ctx = context.WithValue(ctx, "author", "Gentry Rolofson")
		ctx = context.WithValue(ctx, "copyright", "© 2020 bitdev")

		response := r.WithContext(ctx)
		w.Header().Add("x-author", "https://bitdev.io")
		w.Header().Add("siteby", "bitdev")
		w.Header().Add("developer", "Gentry Rolofson")
		w.Header().Add("server", "BlackStar Server")
		w.Header().Add("copyright", "© 2020 BlackStar server by bitdev")
		f, err := os.OpenFile("middleware.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			log.Error(err)

		}
		defer f.Close()
		ctx = r.Context()
		// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
		type ContextKey string
		id := uuid.New()
		var ContextKeyRequestID ContextKey = ContextKey(fmt.Sprintf("requestID-%s", id.String()))
		ctx = context.WithValue(ctx, ContextKeyRequestID, id)
		infoTemp := fasttemplate.New(helpers.GetTemplate("visit_data_complete"), "{{", "}}")

		log.Print("text to append")
		log.Print("more text to append")

		log.Printf("incomming request %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, id.String())
		log.Printf("%s\n", infoTemp.ExecuteString(map[string]interface{}{
			terminal.Colors[25]("time"):       time.Now().Format(time.Stamp),
			terminal.Colors[50]("referer"):    r.Referer(),
			terminal.Colors[100]("HOST"):      r.Host,
			terminal.Colors[200]("Method"):    r.Method,
			terminal.Colors[210]("URI"):       r.RequestURI,
			terminal.Colors[189]("IP"):        r.RemoteAddr,
			terminal.Colors[150]("requestID"): id,
			terminal.Colors[120]("headers"):   r.Header,
			terminal.Colors[108]("Key"):       ContextKeyRequestID,
		}))
		next.ServeHTTP(w, response)
		log.Printf("Finished handling http req. %s", id.String())
	})
}

var handle helpers.HandleLeveled

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(SiteByHeader)
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
	// router.Get("/hospice/eligibility", hospice.Eligibility)
	// router.Get("/hospice/eligibility.html", hospice.Eligibility)
	// router.Get("/hospice/resources", hospice.Resources)
	// router.Get("/hospice/resources.html", hospice.Resources)
	// router.Get("/hospice/community", hospice.Community)
	// router.Get("/hospice/community.html", hospice.Community)
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
			fmt.Println("Static_FS", errors.New("fileServer does not permit URL parameters"))
			log.Error("fileServer does not permit URL parameters")
		}
		http.FileServer(http.Dir("public")).ServeHTTP(w, r)
	})

	return router
}

type server struct {
	Port      string
	IP string
	Address   string
	ENV       *Env
	StartTime time.Time
	Handler *chi.Mux
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

type Env struct {
	Production  bool
	Development bool
}

var color = terminal.NewTerminalColor()
var colors = terminal.Colors
var Mux *chi.Mux

func NewServer(config interface{}) (err error, srv *server) {
	switch c := config.(type) {
	case map[string]interface{}:
		config, ok := config.(map[string]interface{})
		if !ok {
			log.Print(color.RedBlink("Config loaded incorrectly please look at configuration file for errors"))
			return errors.New("config loaded incorrectly please look at configuration file for errors"), nil
		}
		address, ok := config["address"].(string)
		if !ok {
			log.Panic(address)
			return errors.New("failed to load address"), nil
		}
		handler, ok := config["handler"].(*chi.Mux)
		if !ok {
			log.Panic(handler)
			return errors.New("failed to load handler"), nil
		}
		readTimeout, ok := config["readTimeout"].(time.Duration)
		if !ok {
			log.Panic(readTimeout)
			return errors.New("failed to load read timeout"), nil
		}
		readHeaderTimeout, ok := config["readHeaderTimeout"].(time.Duration)
		if !ok {
			log.Panic(readHeaderTimeout)
			return errors.New("failed to load read header timeout"), nil
		}
		writeTimeout, ok := config["writeTimeout"].(time.Duration)
		if !ok {
			log.Print(writeTimeout)
			return errors.New("failed to load write timeout"), nil
		}
		ideaTimeout, ok := config["ideaTimeout"].(time.Duration)
		if !ok {
			log.Print(ideaTimeout)
			return errors.New("failed to load ideaTimeout"), nil
		}
		port, ok := config["port"].(string)
		if !ok {
			log.Print(ideaTimeout)
			return errors.New("failed to load port"), nil
		}
		IP, ok := config["IP"].(string)
		if !ok {
			log.Print(IP)
			return errors.New("failed to load ip"), nil
		}
		ENV, ok := config["env"].(*Env)
		if !ok {
			log.Print(ENV)
			return errors.New("failed to load env"), nil
		}
		startUpTime, ok := config["start_up_time"].(time.Time)
		if !ok {
			log.Print(startUpTime)
			return errors.New("failed to load startUpTime"), nil
		}
		return nil, &server{
			Address: address,
			Port:              port,
			IP: IP,
			ENV: ENV,
			Handler:           handler,
			StartTime: startUpTime,
			ReadTimeout:       readTimeout,
			ReadHeaderTimeout: readHeaderTimeout,
			WriteTimeout:      writeTimeout,
			IdleTimeout:       ideaTimeout,
		}
	case struct {
		Port      string
		IP string
		Address   string
		ENV       *Env
		StartTime time.Time
		Handler *chi.Mux
		ReadTimeout       time.Duration
		ReadHeaderTimeout time.Duration
		WriteTimeout      time.Duration
		IdleTimeout       time.Duration
	}:
		return nil, &server{
			Address: c.Address,
			IP: c.IP,
			Port:   c.Port,
			ENV:      c.ENV,
			StartTime: c.StartTime,
			Handler: c.Handler,
			ReadTimeout:       c.ReadTimeout,
			ReadHeaderTimeout: c.ReadHeaderTimeout,
			WriteTimeout:      c.WriteTimeout,
			IdleTimeout:       c.IdleTimeout,
		}
	case string:
		log.Print("NO STRINGS!")
		return errors.New("strings not supported"), nil
		// make json string side
	case []byte:
		log.Print("NO BYTES!")
		return errors.New("bytes not supported"),nil
		// make json byte style
	default:
		log.Print("FAILING AS HARD AS WE CAN")
		return errors.New("not supported"), nil
	}
}
