package controllers

import (
	"net/http"

	"github.com/xDarkicex/cchha_new_server/helpers"
)

type Hospice Controllers

func (this Hospice) Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r)
}

func (this Hospice) Careers(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) Services(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) Eligibility(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) Community(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) Resources(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) Contact(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) Locations(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}

func (this Hospice) About(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r)
}
