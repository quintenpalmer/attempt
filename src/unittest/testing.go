package unittest

import (
    "testing"
)

type Any interface{}

func Failure(t *testing.T, msg ...Any) {
    t.Fail()
    t.Log(msg)
}

func CheckEqual(t *testing.T, x, y Any) {
    if x != y {
        Failure(t, x, "!=", y)
    }
}

func CheckNotEqual(t *testing.T, x, y Any) {
    if x == y {
        Failure(t, x, "==", y)
    }
}

func Check(t *testing.T, x Any) {
    if x == false {
        Failure(t, x, "== false")
    }
}

func CheckFalse(t *testing.T, x Any) {
    if x == true {
        Failure(t, x, "== true")
    }
}