package helpers

import (
    "fmt"
    "net/http"
    "regexp"
    "strings"
    "time"
    "github.com/alecthomas/template"

    "github.com/xDarkicex/Hospice/terminal"
)

var (
    Splash = 0
    HomeHealth = 1
    Hospice = 2
)

// Execute function renders page with our data
func render(w http.ResponseWriter, r *http.Request, site int, view string, object map[string]interface{}) error {
    handle := NewHandleWithWriter(w)
    device := r.UserAgent()
    expression := regexp.MustCompile("(Mobi(le|/xyz)|Tablet)")
    if !expression.MatchString(device) {
        w.Header().Set("Connection", "keep-alive")
    }
    w.Header().Set("Transfer-Encoding", "gzip, chunked")
    times := make(map[string]interface{})
    times["total"] = time.Now()

    // object["current_user"] = a.User
    object["view"] = view

    funcMap := make(template.FuncMap)
    funcMap["Split"] = func(s string, d string) []string {
        return strings.Split(s, d)
    }
    funcMap["Join"] = func(a []string, b string) string {
        return strings.Join(a, b)
    }
    // funcMap["ParseFlashes"] = func(fucks []interface{}) []Flash {
    // 	var flashes []Flash
    // 	for _, k := range fucks {
    // 		var flash Flash
    // 		json.Unmarshal([]byte(k.(string)), &flash)
    // 		flashes = append(flashes, flash)
    // 	}
    // 	return flashes
    // }
    funcMap["formatPostTime"] = func(t time.Time) string {
        return t.Format(time.UnixDate)
    }

    funcMap["formatTitle"] = func(s string) string {
        title := strings.SplitAfter(s, "/")
        return strings.Title(title[1])
    }

    times["render-page"] = time.Now()
    fmt.Println(terminal.Colors[35](view))
    if site == Hospice {
        gotpl, err := template.New(view).Funcs(funcMap).ParseFiles("./app/views/hospice/layout/navbar.gohtml", "./app/views/hospice/layout/footer.gohtml", "./app/views/hospice/view/"+view+".gohtml", "./app/views/hospice/layout/layout.gohtml")
        if err != nil {
            handle.Error(err)
        }
        err = gotpl.ExecuteTemplate(w, "base", object)
        handle.Error(err)
        times["render-page"] = time.Since(times["render-page"].(time.Time))
        times["total"] = time.Since(times["total"].(time.Time))
        return nil
    }
    if site == Splash {
        gotpl, err := template.New(view).Funcs(funcMap).ParseFiles( "./app/views/splash.gohtml")
        if err != nil {
            handle.Error(err)
        }
        err = gotpl.ExecuteTemplate(w, "base", object)
        handle.Error(err)
        times["render-page"] = time.Since(times["render-page"].(time.Time))
        times["total"] = time.Since(times["total"].(time.Time))
        return nil
    }
    if site == HomeHealth {
        gotpl, err := template.New("base").Funcs(funcMap).ParseFiles( "./app/views/home-health/"+view+".html")
        if err != nil {
            handle.Error(err)
        }
        err = gotpl.ExecuteTemplate(w, "base", object)
        handle.Error(err)
        times["render-page"] = time.Since(times["render-page"].(time.Time))
        times["total"] = time.Since(times["total"].(time.Time))
        return nil
    }
return nil
}

