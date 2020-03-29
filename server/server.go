package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v4"
	"github.com/scorredoira/email"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/xDarkicex/cchha_new_server/app/controllers"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(60 * time.Second))

	// Routes
	// splash
	application := controllers.Application{}
	router.Get("/", application.Index)

	// home health
	homehealth := controllers.HomeHealth{}
	router.Get("/home-health", homehealth.Index)
	router.Get("/home-health/careers", homehealth.Careers)
	router.Get("/home-health/careers.html", homehealth.Careers)
	router.Get("/home-health/services", homehealth.Services)
	router.Get("/home-health/services.html", homehealth.Services)
	router.Get("/home-health/eligibility", homehealth.Eligibility)
	router.Get("/home-health/eligibility.html", homehealth.Eligibility)
	router.Get("/home-health/resources", homehealth.Resources)
	router.Get("/home-health/resources.html", homehealth.Resources)
	router.Get("/home-health/community", homehealth.Community)
	router.Get("/home-health/community.html", homehealth.Community)
	router.Get("/home-health/about", homehealth.About)
	router.Get("/home-health/about.html", homehealth.About)
	router.Get("/home-health/locations", homehealth.Locations)
	router.Get("/home-health/locations.html", homehealth.Locations)
	router.Get("/home-health/contact", homehealth.Contact)
	router.Get("/home-health/contact.html", homehealth.Contact)

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

	router.Post("/contact-careers", func(res http.ResponseWriter, req *http.Request) {
		var cookie *http.Cookie
		name := (req.FormValue("contact_name"))
		e := (req.FormValue("contact_email"))
		add := (req.FormValue("contact_address"))
		phone := (req.FormValue("contact_phone"))
		body := (req.FormValue("contact_body"))
		google_chaptcha := (req.FormValue("g-recaptcha-response"))

		google_struct := validate_google_rechaptcha(google_chaptcha)
		if !google_struct.Success {
			cookie = GenerateCookie("Failed reChaptcha", google_struct.Success)
			http.SetCookie(res, cookie)
			http.Redirect(res, req, "/careers.html", http.StatusFound)
			return
		}
		if !validate_email(e) {
			cookie = GenerateCookie("Must enter valid email", false)
			http.SetCookie(res, cookie)
			http.Redirect(res, req, "/careers.html", http.StatusFound)
			return
		}

		subject := "Contact Request from " + name
		msg := "New Contact request from " + name + "\n" +
			"Address: " + add + "\n" +
			"Email: " + e + "\n" +
			"Phone Number: " + phone + "\n" +
			"Contact Message: " + "\n" +
			body
		m := email.NewMessage(subject, msg)
		m.From = mail.Address{Name: "Jobs", Address: "admin@cchha.com"}
		m.To = []string{"jobs@cchha.com"}
		// HomeHealth2017
		// TODO: encryption example:
		//   e, err := encryption.NewDecryption()
		//   // err handle
		//   e.INI().Decrypt("file path")

		auth := smtp.PlainAuth("", "admin@cchha.com", "Vh2@cchha#G0!", "smtp.gmail.com")
		SMTP := "smtp.gmail.com:587"
		if err := email.Send(SMTP, auth, m); err != nil {
			fmt.Println("Error on send: ")
			fmt.Println(err)
		}
		cookie = GenerateCookie("Email Sent Successful", true)
		http.SetCookie(res, cookie)
		fmt.Println(cookie)
		http.Redirect(res, req, "/", http.StatusFound)
	})

	router.Post("/contact", func(res http.ResponseWriter, req *http.Request) {
		var cookie *http.Cookie
		name := (req.FormValue("contact_name"))
		e := (req.FormValue("contact_email"))
		add := (req.FormValue("contact_address"))
		phone := (req.FormValue("contact_phone"))
		body := (req.FormValue("contact_body"))
		google_chaptcha := (req.FormValue("g-recaptcha-response"))

		google_struct := validate_google_rechaptcha(google_chaptcha)
		if !google_struct.Success {
			cookie = GenerateCookie("Failed reChaptcha", google_struct.Success)
			http.SetCookie(res, cookie)
			http.Redirect(res, req, "/contact.html", http.StatusFound)
			return
		}
		if !validate_email(e) {
			cookie = GenerateCookie("Must enter valid email", false)
			http.SetCookie(res, cookie)
			http.Redirect(res, req, "/contact.html", http.StatusFound)
			return
		}

		subject := "Contact Request from " + name
		msg := "New Contact request from " + name + "\n" +
			"Address: " + add + "\n" +
			"Email: " + e + "\n" +
			"Phone Number: " + phone + "\n" +
			"Contact Message: " + "\n" +
			body
		m := email.NewMessage(subject, msg)
		m.From = mail.Address{Name: "Jobs", Address: "admin@cchha.com"}
		m.To = []string{"info@cchha.com"}
		// HomeHealth2017
		auth := smtp.PlainAuth("", "admin@cchha.com", "Vh2@cchha#G0!", "smtp.gmail.com")
		SMTP := "smtp.gmail.com:587"
		if err := email.Send(SMTP, auth, m); err != nil {
			fmt.Println("Error on send: ")
			fmt.Println(err)
		}
		cookie = GenerateCookie("Email Sent Successful", true)
		http.SetCookie(res, cookie)
		fmt.Println(cookie)
		http.Redirect(res, req, "/", http.StatusFound)
	})

	//static assets
	router.Get("/static/{filepath}*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = chi.URLParam(r, "filepath")
		if strings.ContainsAny(r.URL.Path, "{}*") {
			panic("FileServer does not permit URL parameters.")
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

var mux *chi.Mux

func init() {
	mux = chi.NewRouter()
	fmt.Println("Mux Init, Complete")
}

func validate_email(email string) bool {
	regex, err := regexp.Compile(`\S+@\S+`)
	if err != nil {
		fmt.Println(err)
	}
	if !regex.MatchString(email) {
		return false
	}
	return true
}

func validate_google_rechaptcha(chaptcha string) (r controllers.RecaptchaResponse) {
	google_check := url.Values{
		"secret":   {"6LchUqEUAAAAALM_u_okQofqiw7Htdcp96jJGn1p"},
		"response": {chaptcha},
	}
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", google_check)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	google_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body: %s", err)
	}
	err = json.Unmarshal(google_body, &r)
	if err != nil {
		log.Println("Read error: got invalid JSON: %s", err)
	}
	return r
}

func GenerateCookie(status string, success bool) *http.Cookie {
	type data struct {
		Status  string
		Success bool
	}
	cookie_value := data{
		Status:  status,
		Success: success,
	}
	d, err := json.Marshal(cookie_value)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))
	cookie := &http.Cookie{
		Name:  "message",
		Value: base64.StdEncoding.EncodeToString(d),
		// Path:    "cchha.com/contact.html",
		// Domain:  "cchha.com",
		Expires: time.Now().Add(time.Minute * 1),

		Secure:   false,
		HttpOnly: false,
	}
	// RawExpires: "0",
	// MaxAge:     0,
	return cookie
}
