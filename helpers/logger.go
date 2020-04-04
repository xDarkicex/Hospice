package helpers

import (
    "encoding/json"
    "io"
    log2 "log"

    "github.com/valyala/fasttemplate"

    "github.com/xDarkicex/Hospice/terminal"
)

var LogErrorTemplate string

var LoggerTemplateMap = make(map[string]string)

var (
     // Default *log2.Logger
     //
    Default *log2.Logger
    // Trace *log2.Logger
    //
    Trace   *log2.Logger
    // Info *log2.Logger
    //
    Info    *log2.Logger
    // Warn *log2.Logger
    Warn *log2.Logger
/* ===================
        Error *log2.Logger
     */
    Error   *log2.Logger
)

var Color terminal.Color

var (
    defaultPre = Color.Blue("[" ) + "{{level}}" + Color.Blue("]" ) +" :"
    infoPre = Color.PinkBold("[" ) + "{{level}}" + Color.PinkBold("]" ) +" :"
    tracePre = Color.Orange("[" ) + "{{level}}" + Color.Orange("]" ) +" :"
    errorPre = Color.GreenLight("[" ) + "{{level}}" + Color.GreenLight("]" ) +" :"
    warnPre = Color.Coral("[" ) + "{{level}}" + Color.Coral("]" ) +" :"
)

func Init(defaultHandle io.Writer, infoHandle io.Writer, warnHandle io.Writer, errorHandle io.Writer, traceHandle io.Writer) {
    l := LoggerTemplateMap
    l["date"] = "date={{date}}"
    l["method"] = "method={{Method}}"
    l["host"] = "host={{HOST}}"
    l["uri"] = "uri={{URI}}"
    l["status"] = "status={{status}}"
    l["IP"] = "IP={{IP}}"
    l["referer"] = "referer={{referer}}"
    l["key"] = "key={{key}}"
    l["header"] = "headers={{headers}}"
    l["request_id"] = "request_id={{requestID}}"
    l["error"] = "error={{err}}"
    l["route_info_colored"] = "[{{method}}]" + terminal.Colors[196]("{{host}}{{uri}}")
    l["route info"] = "[{{method}}] {{host}}{{uri}}"
    l["timestamp"] = "{{date}}"
    // not for production use is debugging entry to map.
    l["visit_data_complete"] = l["timestamp"] + l["IP"] + l["referer"] + l["method"] + l["host"] + l["request_id"] + l["key"] + l["uri"] + l["error"] + l["status"]
    rawerrorTemplate := `{ "time": "` + l["timestamp"] + `",` + `"error":` + `"`+l["error"] +`"}`
    // ErrorJSON is fast template json for error logging
    ErrorJSON, err := json.Marshal(rawerrorTemplate)
    l["ErrorJSON"] = string(ErrorJSON)
    Error.Print(err)
    Default = defaultLogger(defaultHandle)
    Error = errorLogger(errorHandle)
    Trace = traceLogger(traceHandle)
    Info = infoLogger(infoHandle)
    Warn = warnLogger(warnHandle)
}

func GetTemplate(pre string) string {
    return LoggerTemplateMap[pre]
}

func defaultLogger(dh io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(defaultPre, "{{","}}", map[string]interface{}{
        "level": Color.PinkLight("Default"),
    })
    defaultPRE := prefixed
    return log2.New(dh, defaultPRE, log2.Ltime|log2.LstdFlags)
}

func errorLogger(eh io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(errorPre, "{{","}}", map[string]interface{}{
        "level":  Color.Red("Error"),
    })
    coloredPRE := prefixed
    return log2.New(eh, coloredPRE, log2.Ldate|log2.Ltime|log2.Lshortfile)
}
func traceLogger(th io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(tracePre, "{{","}}", map[string]interface{}{
        "level": Color.Blue("Trace"),
    })
    coloredPRE := prefixed
    return log2.New(th, coloredPRE, log2.Ldate|log2.Ltime|log2.Lshortfile)
}

func infoLogger(ih io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(infoPre, "{{","}}", map[string]interface{}{
        "level": Color.GreenLight("Info"),
    })
    coloredPRE := prefixed
    return log2.New(ih, coloredPRE, log2.Ldate|log2.Ltime|log2.Lshortfile)
}

func warnLogger(wh io.Writer) *log2.Logger {
    prefixed := fasttemplate.ExecuteString(warnPre, "{{","}}", map[string]interface{}{
        "level": Color.Orange("Warning"),
    })
    return log2.New(wh, prefixed, log2.Ldate|log2.Ltime|log2.Lshortfile)
}

func inti() {

}
