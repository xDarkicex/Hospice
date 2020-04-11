package helpers

import (
    "encoding/json"
    "io"


    "github.com/valyala/fasttemplate"
    log "go.uber.org/zap"

    "github.com/xDarkicex/Hospice/terminal"
)

var LogErrorTemplate string

var LoggerTemplateMap = make(map[string]string)

var (
     // Default *log2.Logger
     //
    Default *log.Logger
    // Trace *log2.Logger
    //
    Trace   *log.Logger
    // Info *log2.Logger
    //
    Info    *log.Logger
    // Warn *log2.Logger
    Warn *log.Logger
/* ===================
        Error *log2.Logger
     */
    Error   *log.Logger
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
    log.Error(err)
    Default = defaultLogger(defaultHandle)
    Error = errorLogger(errorHandle)
    Trace = traceLogger(traceHandle)
    Info = infoLogger(infoHandle)
    Warn = warnLogger(warnHandle)
}

func GetTemplate(pre string) string {
    return LoggerTemplateMap[pre]
}

func defaultLogger(dh io.Writer) *log.Logger {
    // TODO: rewrite the fast templae into a func that takes an error and wraps it in a fasttemplate and returns a string for the logger to use.
     _ = fasttemplate.ExecuteString(defaultPre, "{{","}}", map[string]interface{}{
        "level": Color.PinkLight("Default"),
    })
    logger, err := log.NewDevelopment()
    if err != nil {
        log.Error(err)
    }
    return logger
}

func errorLogger(eh io.Writer) *log.Logger {
    _ = fasttemplate.ExecuteString(errorPre, "{{","}}", map[string]interface{}{
        "level":  Color.Red("Error"),
    })
    logger, err := log.NewDevelopment()
    if err != nil {
        log.Error(err)
    }
    return logger
}
func traceLogger(th io.Writer) *log.Logger {
    _ = fasttemplate.ExecuteString(tracePre, "{{","}}", map[string]interface{}{
        "level": Color.Blue("Trace"),
    })
    logger, err := log.NewDevelopment()
    if err != nil {
        log.Error(err)
    }
    return logger
}

func infoLogger(ih io.Writer) *log.Logger {
    _ = fasttemplate.ExecuteString(infoPre, "{{","}}", map[string]interface{}{
        "level": Color.GreenLight("Info"),
    })
    logger, err := log.NewDevelopment()
    if err != nil {
        log.Error(err)
    }
    return logger
}

func warnLogger(wh io.Writer) *log.Logger {
    _ = fasttemplate.ExecuteString(warnPre, "{{","}}", map[string]interface{}{
        "level": Color.Orange("Warning"),
    })
    logger, err := log.NewDevelopment()
    if err != nil {
        log.Error(err)
    }
    return logger
}
