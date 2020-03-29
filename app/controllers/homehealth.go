package controllers

import (
	"net/http"

	"github.com/xDarkicex/cchha_new_server/helpers"
)

type HomeHealth Controllers

func (this HomeHealth) Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r)
}

func (this HomeHealth) Careers(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) Services(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) Eligibility(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) Community(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) Resources(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) Contact(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) Locations(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this HomeHealth) About(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}
