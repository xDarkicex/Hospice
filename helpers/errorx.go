package helpers

import (
    "context"

    "fmt"
    "net/http"




    _ "github.com/labstack/gommon"
    "github.com/labstack/gommon/log"
   _ "github.com/valyala/fasttemplate"
    "golang.org/x/xerrors"

)


var (
    // DEBUG useful debugging handing of error
    DEBUG level = level{0}
    // INFO minor error or miss in control flow gives print out to terminal
    INFO level = level{1}
    // WARN default error level
    WARN level = level{2}
    // DANGER serious potential for run time error
    DANGER level = level{3}
    // PANIC should fix this dangerous or experimental code in production
    PANIC level = level{4}
)

type Handle struct{
    W http.ResponseWriter
    C context.Context
    E error
    L *log.Logger // std pkg logger support
}

type HandleLeveled struct {
    W http.ResponseWriter
    E error
    Es []error
    C context.Context
    Cache *Cache
    *level
    L *log.Logger // std pkg logger support
}

var (
    pre     = "[{{lvl}}]: "
    route   = "http://{{host}}"
    errLine = "[{{line}}]: {{why}}"
    stdTemp = route + "\n" + errLine
)

// TestErr Simply for testing
var TestErr error

func init() {
    TestErr = xerrors.New("test error handler")
}

func (h *Handle) WithLOGGING(l *log.Logger) *Handle {
    h.L = l
    return h
}

func (h *HandleLeveled) WithLOGGING(l *log.Logger) *HandleLeveled {
    h.L = l
    return h
}

func NewHandle() *Handle {
    return &Handle{W: nil}
}

func NewHandleWithWriter(w http.ResponseWriter) *Handle {
    return &Handle{W: w}
}

// NewHandleLeveled defaults level to 1
func NewHandleLeveled() *HandleLeveled {
    return &HandleLeveled{level:&level{1}}
}

func NewHandleLeveledWithWriter(w http.ResponseWriter, lvl level) *HandleLeveled {
    return &HandleLeveled{W: w, level: &lvl}
}

func NewHandleWithContext(ctx context.Context) *Handle {
    var err error
    if err != nil {
        fmt.Println(">>> NOT NIL <<<")
    }
    return &Handle{
        E: err,
        C:  context.WithValue(ctx, "error", err),
    }
}

func NewHandleLeveledWithContext(ctx context.Context, lvl *level) *HandleLeveled {
    var err error
    if err != nil {
        fmt.Println(">>> NOT NIL  <<<")
    }
    return &HandleLeveled{
        E: err,
        C: ctx,
        level: lvl,
    }
}

func NewHandleLeveledWithCache(cache *Cache, lvl *level) *HandleLeveled {
    var err error
    if err != nil {
        // debugging purposes
        return nil
    }
    return &HandleLeveled{
        E: err,
        C: nil,
        Cache: cache,
        level: lvl,
    }
}


func (h *HandleLeveled) CacheError(key string, err error) {
    h.Cache.Set(key, err)
}
func (h *HandleLeveled) ErrorsFromCache() {
    for _, value := range h.Cache.Object {
        log.Error(value)
    }
}

func (h *HandleLeveled) SetLevel(lvl level) *HandleLeveled {
    h.level = &lvl
    return h
}

// Error I typically use to deal with error inline
// EXAMPLE:
// handle.Error(render(w, r, "home-health/index", object))
func (h *Handle) Error(err error) {
    if err != nil {
        h.E = err
        if h.L != nil {
            pre := h.L.Prefix()
            h.L.Print(pre, stdTemp)
        } else {
            fmt.Println("simple error: ", xerrors.Unwrap(err))
        }
        http.Error(h.W, err.Error(), http.StatusNotFound)
        recover()
    }
    return
}

func (h *Handle) ErrorWithContext(err error) context.Context {
    h.E = err
    if h.L != nil {
        pre :=  h.L.Prefix()
        h.L.Print(pre, stdTemp)
    } else {
        fmt.Println("simple error: ", xerrors.Unwrap(err))
    }
    // return context with error value so you can deal with error where ever you want in program that access context
    return h.C
}

func (h *HandleLeveled) Error(err error) *HandleLeveled {
    h.E = err
    h.L.Error(err)
    return h
}
