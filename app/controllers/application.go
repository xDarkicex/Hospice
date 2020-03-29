package controllers

import (
	"net/http"

	"github.com/xDarkicex/cchha_new_server/helpers"
)

type Application Controllers
// TODO: Stage one
//  Make Error Handler
//  Update Render function
//  Implement templating
func (this Application) Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r)
}
