package controllers

import (
	"net/http"

	"github.com/xDarkicex/cchha_new_server/helpers"
)

type HomeHealth Controllers

func (this HomeHealth) Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r, "hospice/index", map[string]interface{}{
		"Title": "Hospice Home Page",
		"View": "index",
	})
}

func (this HomeHealth) Careers(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/careers",  map[string]interface{}{
		"Title": "Hospice Careers",
		"View": "careers",
	})
}

func (this HomeHealth) Services(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/services", map[string]interface{}{
		"Title": "Hospice Services",
		"View": "services",
	})
}

func (this HomeHealth) Eligibility(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/index", map[string]interface{}{
		"Title": "Hospice Eligibility",
		"View": "eligibility",
	})
}

func (this HomeHealth) Community(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/index", map[string]interface{}{
		"Title": "Hospice Community",
		"View": "community",
	})
}

func (this HomeHealth) Resources(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/index", map[string]interface{}{
		"Title": "Hospice Resources",
		"View": "resources",
	})
}

func (this HomeHealth) Contact(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/Contact", map[string]interface{}{
		"Title": "Hospice Contact",
		"View": "contact",
	})
}

func (this HomeHealth) Locations(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/locations", map[string]interface{}{
		"Title": "Hospice Locations",
		"View": "locations",
	})
}

func (this HomeHealth) About(w http.ResponseWriter, r *http.Request) {
	helpers.RedirectWithoutHTML(w, r)
	helpers.Render(w, r, "hospice/about", map[string]interface{}{
		"Title": "Hospice About",
		"View": "about",
	})
}

