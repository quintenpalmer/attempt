package gameserver

import (
    "errors"
    "regexp"
)

/*
Manages url regular expressions and their associated callback functions
for the webserver. Direct access is hidden from the users of the webserver
*/

type regexpCallback struct {
    pattern *regexp.Regexp
    callback interface{}
}


type callbackHandler struct {
    patterns map[string]regexpCallback
}

// Should not be exposed
func newcallbackHandler() *callbackHandler {
    cbh := new(callbackHandler)
    cbh.patterns = make(map[string]regexpCallback)
    return cbh
}

// Should be exposed, not tied to urls, probably to Webserver
func (cbh *callbackHandler) registerCallback(url string, callback interface{}) error {
    rexp, error := regexp.Compile(url)
    if (error != nil) {
        return error
    }
    rexpcall := regexpCallback{rexp, callback}

    cbh.patterns[url] = rexpcall

    return nil
}

func (cbh *callbackHandler) findCallback(url string) (interface{}, []string, error) {
    for _, v := range cbh.patterns {
        res := v.pattern.FindStringSubmatch(url)
        if res != nil {
            return v.callback, res, nil
        }
    }

    return nil, nil, errors.New("Callback matching " + url + " does not exist")
}
