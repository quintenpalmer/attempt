package engine

import (
    "time"
)

func Timer(waitTime time.Duration, thunk func()) {
    go func() {
        time.Sleep(waitTime)
        thunk()
    }()
}

func RepeatingTimer(waitTime time.Duration, boolFunc func() bool) {
    go func () {
        for {
            time.Sleep(waitTime)
            if boolFunc() {
                return
            }
        }
    }()
}
