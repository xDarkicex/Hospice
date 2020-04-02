package controllers

import (
	"net/http"

	"github.com/xDarkicex/Hospice/helpers"
)

type HomeHealth Controllers
func (this HomeHealth) Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r, helpers.HomeHealth,"index", map[string]interface{}{
		"Title": "Hospice Home Page",
		"View": "index",
	})
}

func (this HomeHealth) Careers(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"careers",  map[string]interface{}{
		"Title": "Hospice Careers",
		"View": "careers",
	})
}

func (this HomeHealth) Services(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"services", map[string]interface{}{
		"Title": "Hospice Services",
		"View": "services",
	})
}

func (this HomeHealth) Eligibility(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"eligibility", map[string]interface{}{
		"Title": "Hospice Eligibility",
		"View": "eligibility",
	})
}

func (this HomeHealth) Community(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"community", map[string]interface{}{
		"Title": "Hospice Community",
		"View": "community",
	})
}

func (this HomeHealth) Resources(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"resources", map[string]interface{}{
		"Title": "Hospice Resources",
		"View": "resources",
	})
}

func (this HomeHealth) Contact(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"contact", map[string]interface{}{
		"Title": "Hospice Contact",
		"View": "contact",
	})
}

func (this HomeHealth) Locations(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r, helpers.HomeHealth,"locations", map[string]interface{}{
		"Title": "Hospice Locations",
		"View": "locations",
	})
}

func (this HomeHealth) About(w http.ResponseWriter, r *http.Request) {
	// handler.Error(helpers.RedirectWithoutHTML(w, r))
	helpers.Render(w, r,helpers.HomeHealth, "about", map[string]interface{}{
		"Title": "Hospice About",
		"View": "about",
	})
}

