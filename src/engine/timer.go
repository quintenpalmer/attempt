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