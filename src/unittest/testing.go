package unittest

import (
    "testing"
    "reflect"
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

func CheckNil(t *testing.T, x Any) {
    if x == nil {
        Failure(t, x, "!= nil")
    }
}

func CheckNotNil(t *testing.T, x Any) {
    if x == nil {
        Failure(t, x, "== nil")
    }
}

func CheckDeepEqual(t *testing.T, x, y Any) {
    if !reflect.DeepEqual(x, y) {
        Failure(t, x, "!=", y)
    }
}

func CheckDeepNotEqual(t *testing.T, x, y Any) {
    if reflect.DeepEqual(x, y) {
        Failure(t, x, "==", y)
    }
}
