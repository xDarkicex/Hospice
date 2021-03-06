package controllers

import (
	"net/http"

	"github.com/xDarkicex/Hospice/helpers"
)

type Hospice Controllers

var hospice = helpers.Hospice
func (this Hospice) Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r, hospice, "index", map[string]interface{}{
		"Title": "Hospice Home Page",
		"View": "index",
	})
}

func (this Hospice) Careers(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice, "careers",  map[string]interface{}{
		"Title": "Hospice Careers",
		"View": "careers",
	})
}

func (this Hospice) Services(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"services", map[string]interface{}{
		"Title": "Hospice Services",
		"View": "services",
	})
}

func (this Hospice) Eligibility(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"eligibility", map[string]interface{}{
		"Title": "Hospice Eligibility",
		"View": "eligibility",
	})
}

func (this Hospice) Community(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"community", map[string]interface{}{
		"Title": "Hospice Community",
		"View": "community",
	})
}

func (this Hospice) Resources(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"resources", map[string]interface{}{
		"Title": "Hospice Resources",
		"View": "resources",
	})
}

func (this Hospice) Contact(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"contact", map[string]interface{}{
		"Title": "Hospice Contact",
		"View": "contact",
	})
}

func (this Hospice) Locations(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"locations", map[string]interface{}{
		"Title": "Hospice Locations",
		"View": "locations",
	})
}

func (this Hospice) About(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, hospice,"about", map[string]interface{}{
		"Title": "Hospice About",
		"View": "about",
	})
}
