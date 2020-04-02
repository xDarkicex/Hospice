package controllers

import (
	"net/http"
	_ "regexp"
	_ "strings"
	_ "time"

	"github.com/xDarkicex/Hospice/helpers"
)

// Application  binding type
type Application Controllers
// TODO: Stage one
//  Make Error Handler
//  Update Render function
//  Implement templating
func (app Application) Index(w http.ResponseWriter, r *http.Request) {

	// helpers.Render(w, r, "splash", map[string]interface{}{"Index": "Home Page"})
	// helpers.RenderSplash(w, r, "splash", map[string]interface{}{"Index": "Home Page"})
	helpers.Render(w, r, helpers.Splash, "/",map[string]interface{}{
		"Title":     "Home Page",
		"View":      "index",
		"Version":   "v1.9.5",
		"Author":    "Gentry Rolofson",
		"Copyright": "© 2020 Compassionate Care Hospice Central California",
		"Siteby":    "https://bitdev.io",
	})
}

