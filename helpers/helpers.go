package helpers

import (
	"fmt"
	"net/http"
	"strings"
)

func RedirectWithoutHTML(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path == "/" {
		return nil
	}
	if !strings.Contains(r.URL.Path, ".") {
		// force default all non-specific paths to .html
		http.Redirect(w, r, r.URL.Path+".html", http.StatusFound)
		return nil
	}
	return nil
}

func withoutHTML(w http.ResponseWriter, r *http.Request) string {
	if strings.Contains(r.URL.EscapedPath(), ".") {
		fmt.Println(r.URL.Path)
		path := strings.Split(r.URL.EscapedPath(), ".")
		return path[0]
	}
	return r.URL.EscapedPath()
}

func Render(w http.ResponseWriter, r *http.Request, view string, object map[string]interface{}) {
	handle := NewHandleWithWriter(w)
	path := withoutHTML(w, r)
	if path == "/" {
		handle.Error(render(w,r, "splash", object))
		return
	}
	if path == "/home-health" {
		handle.Error(render(w, r, "home-health/index", object))
		return
	}
	if path == "/hospice" {
		handle.Error(render(w,r,"hospice/index", object))
		return
	}
	handle.Error(render(w,r, path, object))
}
