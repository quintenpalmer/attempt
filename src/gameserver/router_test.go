package gameserver

import (
    "net/http"
    "testing"
)

func TestRouteValidRequest(t *testing.T) {
    ro := newrouter()
    ro.registerCallback("^/test/([\\w]+)$", routerTestFunc)
    
    hr, _ := http.NewRequest("GET", "http://localhost:8000/test/waffle", nil)
    res := ro.routeRequest(hr)
    if res.Body != "waffle" {
        t.Fail()
    }
    if res.Status != 200 {
        t.Fail()
    }
}

func TestRouteInvalidRequest(t *testing.T) {
    ro := newrouter()
    ro.registerCallback("^/test/([\\w]+)$", routerTestFunc)
    
    hr, _ := http.NewRequest("GET", "http://localhost:8000/test/waffle/cheese", nil)
    res := ro.routeRequest(hr)
    if res.Body != "No handler for path: /test/waffle/cheese" {
        t.Fail()
    }
    if res.Status != 404 {
        t.Fail()
    }
}

func routerTestFunc(r *http.Request, a string) HttpResponse {
    return HttpResponse{a, 200}
}
