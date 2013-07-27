package gameserver

import (    
    "testing"
    "reflect"
)

func TestRegisterValidCallback(t *testing.T) {
    x := newcallbackHandler()
    e := x.registerCallback("test", urlsTestFunc)
    if e != nil {
        t.Fail()
    }
}

func TestRegisterInvalidCallback(t *testing.T) {
    x := newcallbackHandler()
    e := x.registerCallback("test[\\w", urlsTestFunc)
    if e == nil {
        t.Fail()
    }
}

func TestFindExistingCallback(t *testing.T) {
    x := newcallbackHandler()
    e := x.registerCallback("test/([\\w]+)/", urlsTestFunc)
    if e != nil {
        t.Fail()
    }
    cb, args, er := x.findCallback("test/waffle/")
    if cb == nil {
        t.Fail()
    }
    if args == nil {
        t.Fail()
    }
    if er != nil {
        t.Fail()
    }
    if reflect.ValueOf(cb) != reflect.ValueOf(urlsTestFunc) {
        t.Fail()
    }
    // The first argument is the URL string itself
    if args[1] != "waffle" {
        t.Fail()
    }
}

func TestFindNonExistingCallback(t *testing.T) {
    x := newcallbackHandler()
    e := x.registerCallback("^test/[\\w]/$", urlsTestFunc)
    if e != nil {
        t.Fail()
    }
    cb, args, er := x.findCallback("waffle/")
    if cb != nil {
        t.Fail()
    }
    if args != nil {
        t.Fail()
    }
    if er == nil {
        t.Fail()
    }
}

func urlsTestFunc() {

}
