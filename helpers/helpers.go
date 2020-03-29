package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	// if r.URL.Path == "/" {
	// 	return nil
	// }
	// fmt.Println(r.URL.Path)
	if strings.Contains(r.URL.EscapedPath(), ".") {
		fmt.Println(r.URL.Path)
		path := strings.Split(r.URL.EscapedPath(), ".")
		return path[0]
	}
	return r.URL.EscapedPath()
}

func Render(w http.ResponseWriter, r *http.Request) {
	path := withoutHTML(w, r)
	device := r.UserAgent()
	expression := regexp.MustCompile("(Mobi(le|/xyz)|Tablet)")
	if !expression.MatchString(device) {
		w.Header().Set("Connection", "keep-alive")
	}
	if path == "/" {
		file, err := ioutil.ReadFile("app/views/splash.html")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(file))
		return
	}
	if path == "/home-health" {
		file, err := ioutil.ReadFile("app/views" + path + "/index.html")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(file))
		return
	}
	if path == "/hospice" {
		file, err := ioutil.ReadFile("app/views/" + path + "/index.html")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(file))
		return
	}
	file, err := ioutil.ReadFile("app/views/" + path + ".html")
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(file))
}
