package gameserver

import (
    "fmt"
    "net/http"
    "reflect"
    "regexp"
    "strconv"
)

/*
Interacts with urls.go to route requests to their corresponding functions.
Wires the capture groups from a url's regex to the callbacks arguments
*/

var allChars = regexp.MustCompile("[\\w]+")
var allNum = regexp.MustCompile("[\\d]+")

type router struct {
    cbh *callbackHandler
}

func newrouter() *router {
    cbh := newcallbackHandler()
    r := router{cbh}

    return &r
}

func (r *router) registerCallback(url string, callback interface{}) error {
    err := r.cbh.registerCallback(url, callback)

    return err
}

func (r *router) routeRequest(req *http.Request) HttpResponse {
    fmt.Println(req.URL.Path)
    cb, subs, err := r.cbh.findCallback(req.URL.Path)
    if err != nil {
        return HttpResponse{"No handler for path: " + req.URL.Path, 404}
    }

    fmt.Println(subs)
    cbV := reflect.ValueOf(cb)
    fmt.Println(reflect.TypeOf(cb))
    // The first arg is the string itself if there's a match
    args := convertToReflectValues(subs)
    args[0] = reflect.ValueOf(req)
    
    ret := cbV.Call(args)
    return reconstructHttpResponse(ret[0])
}

func convertToReflectValues(args []string) []reflect.Value {
    o := make([]reflect.Value, len(args))
    for i := range args {
        if allNum.MatchString(args[i]) {
            // we already know it matches only numbers so we can ignore the error
            cv, _ := strconv.Atoi(args[i])
            o[i] = reflect.ValueOf(cv)
        }else {
            o[i] = reflect.ValueOf(args[i])
        }
    }
    return o
}

func reconstructHttpResponse(val reflect.Value) HttpResponse {
    body := val.Field(0).String()
    status := val.Field(1).Int()

    return HttpResponse{body, int(status)}
}
