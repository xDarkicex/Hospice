package helpers

import (
    "errors"
    log2 "log"
    "sync"
    "time"
)
type level struct {
    int
}

var (
    Rapid = level{0}
    Short = level{1}
    Normal = level{2}
    Long = level{3}
)

const rapidExpiration = time.Second * time.Duration(30) // testing
const shortExpiration = time.Minute * time.Duration(30)
const defaultExpiration = time.Hour * time.Duration(1)
const longExpiration = time.Hour * time.Duration(2) // deep freeze
type Cache struct {
    Object map[string]*Object
    mutex *sync.Mutex
}
type Object struct {
    LEVEL int
    Expire time.Time
    Object interface{}
}

func (l level) setExpires() time.Duration {
    switch l.int {
        case 0:
            log2.Print("level set to rapid expire")
            return rapidExpiration
        case 1:
            log2.Print("level set to short expire")
            return shortExpiration
        case 2:
            log2.Print("level set to normal expire")
            return defaultExpiration
        case 3:
            log2.Print("level set to long expire")
            return longExpiration
    default:
        log2.Print("level set error default now")
        return time.Second * 0
    }

}
// Get cached Object
// key: the key our cached object.
// Example:
// object := Cache.Get("layout")
var cache = new(Cache)
func init() {

    var user = struct{
        Name string
        Age int
    }{
        Name: "Gentry",
        Age: 27,
    }

    cache.Set("current_User", user)
    cache.Get("Current_user")
}

func(cache *Cache) Get(key string) (err error, c *Object) {
    // cache.mutex.Lock()
    if cache.Object[key] == nil {
        return errors.New("cached object no longer exists in memory"), nil
    }
    // cache.mutex.Unlock()
    return nil, cache.Object[key]
}

func(cache *Cache) Set(key string, value interface{}) (err error, c *Object) {
    // cache.mutex.Lock()
    if cache.Object[key] != nil {
        return errors.New("cached object already exists in memory"), cache.Object[key]
    }
    lvl := level{int: 3}
    object := NewObject(time.Now().Add(lvl.setExpires()), value)
    object.LEVEL = 3
    // cache.mutex.Unlock()
    return nil, object
}

// DeleteCache object for updating pages
func (cache *Cache) Remove(key string) {
    delete(cache.Object, key)
}

// NewCache Creates Pointer to new cache can make many caches
func NewCache() *Cache {
    return &Cache{}
}

func NewObject(exp time.Time, object interface{}) *Object {
    return &Object{
        Expire: exp,
        Object: object,
    }
}

// func init() {
//     var timer *time.Timer
//     timer = time.AfterFunc(defaultExpiration, func() {
//         for key, cachedObject := range cache {
//             if time.Now().After(cachedObject.Expire) {
//                 delete(cache, key)
//             }
//         }
//         timer.Reset(defaultExpiration)
//     })
//
// }
