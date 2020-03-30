package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/url"
	"regexp"
	"time"

	"github.com/scorredoira/email"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type Mailer Controllers






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

func validate_google_rechaptcha(chaptcha string) (r RecaptchaResponse) {
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


func (mailer *Mailer) ContactCareers(res http.ResponseWriter, req *http.Request) {
	var cookie *http.Cookie
	name := req.FormValue("contact_name")
	e := req.FormValue("contact_email")
	add := req.FormValue("contact_address")
	phone := req.FormValue("contact_phone")
	body := req.FormValue("contact_body")
	google_chaptcha := req.FormValue("g-recaptcha-response")

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
}

func(mailer *Mailer) Contact(res http.ResponseWriter, req *http.Request) {
	var cookie *http.Cookie
	name := req.FormValue("contact_name")
	e := req.FormValue("contact_email")
	add := req.FormValue("contact_address")
	phone := req.FormValue("contact_phone")
	body := req.FormValue("contact_body")
	google_chaptcha := req.FormValue("g-recaptcha-response")

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
}
