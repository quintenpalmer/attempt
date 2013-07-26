package gogame

import (
    "testing"
)

func Failure(t *testing.T, x, y Any, mid string) {
    t.Fail()
    t.Log(x, mid, y)
}

func CheckEqual(t *testing.T, x, y Any) {
    if x != y {
        Failure(t, x, y, "!=")
    }
}

func CheckNotEqual(t *testing.T, x, y Any) {
    if x == y {
        Failure(t, x, y, "==")
    }
}